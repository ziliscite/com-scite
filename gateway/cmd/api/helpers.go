package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"strings"
)

var (
	ErrBadlyFormattedJSON = errors.New("body contains badly-formed JSON")
	ErrInvalidTypeJSON    = errors.New("body contains incorrect JSON type")
)

func writeJSON(w http.ResponseWriter, code int, data any, headers ...http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func sendError(w http.ResponseWriter, code int, e error) {
	if err := writeJSON(w, code, map[string]any{
		"error": e.Error(),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func sendGRPCError(w http.ResponseWriter, e error) {
	if st, ok := status.FromError(e); ok {
		switch st.Code() {
		case codes.DeadlineExceeded:
			sendError(w, http.StatusGatewayTimeout, errors.New("request timed out. Please try again later"))
		case codes.Canceled:
			sendError(w, http.StatusBadRequest, errors.New("request was canceled by the client"))
		case codes.Unauthenticated:
			sendError(w, http.StatusUnauthorized, errors.New("authentication failed. Please check your credentials"))
		case codes.PermissionDenied:
			sendError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))

		case codes.InvalidArgument:
			sendError(w, http.StatusBadRequest, errors.New("invalid input provided. Please check your request"))

		case codes.NotFound:
			sendError(w, http.StatusNotFound, errors.New("resource not found"))
		case codes.AlreadyExists:
			sendError(w, http.StatusConflict, errors.New("resource already exists. Please check your request"))
		case codes.FailedPrecondition:
			sendError(w, http.StatusPreconditionFailed, errors.New("precondition failed. Check the request constraints"))
		case codes.Aborted:
			sendError(w, http.StatusConflict, errors.New("request was aborted. Please try again later"))
		case codes.OutOfRange:
			sendError(w, http.StatusRequestedRangeNotSatisfiable, errors.New("the requested range is out of bounds"))
		case codes.Internal:
			sendError(w, http.StatusInternalServerError, errors.New("an internal server error occurred. Please try again later"))
		case codes.Unavailable:
			sendError(w, http.StatusServiceUnavailable, errors.New("service unavailable. Please try again later"))
		case codes.DataLoss:
			sendError(w, http.StatusInternalServerError, errors.New("data loss detected. Please contact support"))
		default:
			sendError(w, http.StatusInternalServerError, errors.New("an unknown error occurred. Please try again later"))
		}
	} else {
		sendError(w, http.StatusInternalServerError, e)
	}
}

func readBody(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("%w (at character %d)", ErrBadlyFormattedJSON, syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return ErrBadlyFormattedJSON

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("%w for field %q", ErrInvalidTypeJSON, unmarshalTypeError.Field)
			}
			return fmt.Errorf("%w (at character %d)", ErrInvalidTypeJSON, unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key: '%s'", strings.Trim(fieldName, "\""))

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
