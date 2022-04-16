package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/blue-health/go-backend-exercise/app/storage"
	"github.com/blue-health/go-backend-exercise/app/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kelseyhightower/envconfig"
)

// Any runtime configuration your app will require should
// be placed here. This struct defines the environmental variables
// the program will look for. For instance, the database address will
// be parsed from the GO_EXERCISE_DATABASE_DSN environmental variable.
type config struct {
	Server struct {
		Address         string        `default:":8080"`
		ReadTimeout     time.Duration `split_words:"true" default:"5s"`
		WriteTimeout    time.Duration `split_words:"true" default:"10s"`
		IdleTimeout     time.Duration `split_words:"true" default:"15s"`
		ShutdownTimeout time.Duration `split_words:"true" default:"30s"`
		RequestTimeout  time.Duration `split_words:"true" default:"45s"`
	}
	Database struct {
		DSN string `required:"true"`
	}
}

const app = "go_exercise"

func main() {
	ctx := context.Background()

	var cfg config
	if err := envconfig.Process(app, &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := run(ctx, &cfg); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}

func run(ctx context.Context, cfg *config) error {
	db, err := storage.NewPostgres(ctx, cfg.Database.DSN)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatalf("failed to close the database: %v", err)
		}
	}()

	var (
		done   = make(chan struct{})
		quit   = make(chan os.Signal, 1)
		server = newServer(cfg, cfg.Server.Address, func(m chi.Router) {
			m.Mount("/accounts", web.NewAccountService(nil).Router())
		})
	)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit

		log.Println("server is shutting down...")

		sctx, cancel := context.WithTimeout(ctx, cfg.Server.ShutdownTimeout)
		defer cancel()

		server.SetKeepAlivesEnabled(false)

		if err := server.Shutdown(sctx); err != nil {
			log.Fatalf("failed to gracefully shutdown server: %v", err)
		}

		close(done)
	}()

	log.Println("starting server at", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server at %s: %w", server.Addr, err)
	}

	<-done

	log.Println("server stopped")

	return nil
}

func newServer(cfg *config, address string, f func(chi.Router)) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(cfg.Server.RequestTimeout))

	f(r)

	return &http.Server{
		Addr:         address,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		Handler:      r,
	}
}
