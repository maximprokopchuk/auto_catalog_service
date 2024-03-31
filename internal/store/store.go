package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/maximprokopchuk/auto_catalog_service/internal/sqlc"
)

type Store struct {
	config     *Config
	Connection *pgx.Conn
	Queries    *sqlc.Queries
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open(ctx context.Context) error {
	url := fmt.Sprintf(
		"host=%s dbname=%s sslmode=disable user=%s password=%s",
		s.config.DatabaseHost,
		s.config.DatabaseName,
		s.config.DatabaseUser,
		s.config.DatabasePassword,
	)

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}
	queries := sqlc.New(conn)

	s.Connection = conn
	s.Queries = queries

	return nil
}

func (s *Store) Close(ctx context.Context) {
	defer s.Connection.Close(ctx)
}
