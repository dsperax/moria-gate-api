package ports

import (
	"context"
	"moria-gate-api/internal/core/domain"
)

type ClientRepository interface {
	FindByDocument(ctx context.Context, doc string) (*domain.Client, error)
}
