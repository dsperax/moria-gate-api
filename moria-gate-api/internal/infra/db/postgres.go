package db

import (
	"context"
	"database/sql"
	"fmt"

	"moria-gate-api/internal/core/domain"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	return &PostgresRepository{DB: db}, nil
}

func (r *PostgresRepository) FindByDocument(ctx context.Context, doc string) (*domain.Client, error) {
	var c domain.Client
	var a domain.Address
	var cs domain.CreditInfo

	// Join tables based on the document (used as foreign key)
	err := r.DB.QueryRowContext(ctx, `
		SELECT c.name, c.document, 
		       a.street, a.city, a.state, a.district, a.zip, a.country,
		       s.situation, s.risk, s.type
		FROM clients c
		JOIN addresses a ON c.document = a.document
		JOIN credit_status s ON c.document = s.document
		WHERE c.document = $1
	`, doc).Scan(&c.Name, &c.Document,
		&a.Street, &a.City, &a.State, &a.District, &a.Zip, &a.Country,
		&cs.Situation, &cs.Risk, &cs.Type)

	if err != nil {
		return nil, fmt.Errorf("not found or error: %w", err)
	}

	c.Address = a
	c.Credit = cs

	return &c, nil
}
