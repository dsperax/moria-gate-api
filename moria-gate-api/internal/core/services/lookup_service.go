package services

import (
	"context"
	"errors"
	"moria-gate-api/internal/core/domain"
	"moria-gate-api/internal/core/ports"
)

type LookupService struct {
	repo ports.ClientRepository
}

func NewLookupService(r ports.ClientRepository) *LookupService {
	return &LookupService{repo: r}
}

func (s *LookupService) GetSecureClientData(ctx context.Context, doc domain.Document) (*domain.Client, error) {
	if !doc.IsValid() {
		return nil, errors.New("invalid document format")
	}
	return s.repo.FindByDocument(ctx, doc.Key())
}
