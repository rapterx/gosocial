package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rapterx/GoSocial/internal/store"
)

	type application struct {
		config config
		store store.Storage
	}

	type config struct {
		addr string
		db    dbConfig
		env string
	}

	type dbConfig struct {
		addr string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime string
	}

	func (app *application) mount() http.Handler {
		r := chi.NewRouter()
		r.Get("/health", app.healthCheckHandler)
		
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)

		return r
		
	}

	func (app *application) run(mux http.Handler) error {
		srv := &http.Server{
			Addr: app.config.addr,
			Handler: mux,
			WriteTimeout: time.Second * 30,
			ReadTimeout: time.Second * 10,
			IdleTimeout: time.Minute,
		}

		log.Printf("Server started at %s", app.config.addr)

		return srv.ListenAndServe()
	}