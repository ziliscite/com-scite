package main

import (
	"bufio"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	pb "github.com/ziliscite/micro-auth/gateway/pkg/protobuf"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type applications struct {
	auc pb.AuthServiceClient
	atc pb.ActivationServiceClient
	cc  pb.ComicServiceClient
	cvc pb.CoverServiceClient
}

func (app *applications) register(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readBody(w, r, &requestBody)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := app.auc.Register(r.Context(), &pb.RegisterRequest{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, resp); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *applications) activate(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		TokenString string `json:"token_string"`
	}

	err := readBody(w, r, &requestBody)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := app.atc.ActivateUser(r.Context(), &pb.ActivateRequest{
		TokenString: requestBody.TokenString,
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, resp); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *applications) login(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readBody(w, r, &requestBody)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := app.auc.Login(r.Context(), &pb.LoginRequest{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, resp); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *applications) newComic(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var requestBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Artist      string `json:"artist"`

		Status string `json:"status"`
		Type   string `json:"type"`

		Genres []string `json:"genres"`
	}
	if err := readBody(w, r, &requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	comic, err := app.cc.InsertComic(ctx, &pb.InsertComicRequest{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Author:      requestBody.Author,
		Artist:      requestBody.Artist,

		Status: requestBody.Status,
		Type:   requestBody.Type,

		Genres: requestBody.Genres,

		// Kudu streaming data multipart form buat upload cover
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, comic); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *applications) getComicBySlug(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	slug := r.PathValue("slug")

	comic, err := app.cc.GetComicBySlug(ctx, &pb.GetComicBySlugRequest{
		Slug: slug,
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, comic); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

const maxImageSize = 10 << 20

func (app *applications) newCover(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	slug := r.PathValue("slug")

	comic, err := app.cc.GetComicBySlug(ctx, &pb.GetComicBySlugRequest{
		Slug: slug,
	})
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxImageSize)

	// Parse the multipart form
	if err = r.ParseMultipartForm(maxImageSize); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	// Retrieve the file from the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	stream, err := app.cvc.UploadCover(ctx)
	if err != nil {
		sendGRPCError(w, err)
		return
	}

	// build metadata to be sent first
	req := &pb.UploadCoverRequest{
		Data: &pb.UploadCoverRequest_Metadata{
			Metadata: &pb.CoverMetadata{
				Filename: header.Filename,
				ComicId:  comic.Comic.Id,
			},
		},
	}

	// send the first request to the server
	if err = stream.Send(req); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	// create buffer to send a file
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		// read the data to the buffer
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}

		req := &pb.UploadCoverRequest{
			Data: &pb.UploadCoverRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		// send chunk to the server
		if err = stream.Send(req); err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, res); err != nil {
		slog.Error("failed to write json", "error", err.Error())
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *applications) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(RateLimit)
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/v0/register", app.register)
	r.Post("/v0/activate", app.activate)
	r.Post("/v0/login", app.login)

	r.Post("/v0/comic/{slug}/cover", app.newCover)
	r.Get("/v0/comic/{slug}", app.getComicBySlug)
	r.Post("/v0/comic", app.newComic)

	return r
}
