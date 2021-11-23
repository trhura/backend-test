package main

import (
	"hlaing-backend-test-1/api"
	"hlaing-backend-test-1/dbm"
	"hlaing-backend-test-1/env"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	envars := env.GetEnvironment()
	dbManager := dbm.NewDatabaseManager(envars)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/users", api.NewAPIRouter(dbManager))
	http.ListenAndServe(":3000", r)
}
