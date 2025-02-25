package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	pb "github.com/ziliscite/micro-auth/gateway/pkg/protobuf"
	"net/http"
	"time"
)

type applications struct {
	auc pb.AuthServiceClient
	atc pb.ActivationServiceClient
	cc  pb.ComicServiceClient
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

	r.Post("/v0/comic", app.newComic)
	r.Post("/v0/comic/{slug}", app.getComicBySlug)

	return r
}
