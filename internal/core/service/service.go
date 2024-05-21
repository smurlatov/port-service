package service

import (
	"context"
	"port-service/internal/core/domain"
)

type PortRepository interface {
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
}

type PortService struct {
	repo PortRepository
}

func NewPortService(repo PortRepository) *PortService {
	return &PortService{
		repo: repo,
	}
}

func (s PortService) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	return s.repo.CreateOrUpdatePort(ctx, port)
}
