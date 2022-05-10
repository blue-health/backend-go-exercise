package main

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type config struct {
	Up       bool          `default:"true"`
	Delay    time.Duration `default:"1s"`
	Database struct {
		DSN string `required:"true"`
	}
}

const app = "go_exercise_migrate"

//go:embed sql/*.sql
var migrationsFs embed.FS

func main() {
	var cfg config
	if err := envconfig.Process(app, &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	time.Sleep(cfg.Delay)

	if err := run(&cfg); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}

func run(cfg *config) error {
	if cfg.Up {
		log.Println("migrating UP")
	} else {
		log.Println("migrating DOWN")
	}

	d, err := iofs.New(migrationsFs, "sql")
	if err != nil {
		return fmt.Errorf("failed to create iofs: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.Database.DSN)
	if err != nil {
		return fmt.Errorf("failed to create migration source: %w", err)
	}

	f := m.Up

	if !cfg.Up {
		f = m.Down
	}

	if err = f(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
