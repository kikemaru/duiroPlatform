package repository

import (
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kikemaru/duiroPlatform/config"
	"github.com/pressly/goose/v3"
)

// Postgres - .
type Postgres struct {
	*sqlx.DB
}

// New -.
func New(cfg *config.Db) (*Postgres, error) {
	db, err := sqlx.Open("pgx", cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	pgdb := Postgres{
		db,
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Minute)

	// migrations
	goose.SetTableName("migrations")

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(db.DB, cfg.MigrationsSourceURL); err != nil {
		return nil, err
	}

	return &pgdb, nil
}
