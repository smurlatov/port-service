package repository

import (
	"context"
	"port-service/internal/core/domain"
	"time"
)

type Port struct {
	Id          string
	Name        string
	Code        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	CreateOrUpdatePort(ctx context.Context, port Port) error
}

type PortRepository struct {
	storage Storage
}

func NewPortRepository(s Storage) *PortRepository {
	return &PortRepository{
		storage: s,
	}
}

func (p PortRepository) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	if port == nil {
		return domain.ErrNil
	}
	err := p.storage.CreateOrUpdatePort(ctx, convertDomainToStore(port))
	if err != nil {
		return err
	}

	return nil
}

func convertDomainToStore(p *domain.Port) Port {
	return Port{
		Id:          p.Id(),
		Name:        p.Name(),
		Code:        p.Code(),
		City:        p.City(),
		Country:     p.Country(),
		Alias:       append([]string(nil), p.Alias()...),
		Regions:     append([]string(nil), p.Regions()...),
		Coordinates: append([]float64(nil), p.Coordinates()...),
		Province:    p.Province(),
		Timezone:    p.Timezone(),
		Unlocs:      append([]string(nil), p.Unlocs()...),
	}
}
