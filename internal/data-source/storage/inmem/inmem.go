package inmem

import (
	"context"
	"port-service/internal/core/domain"
	"port-service/internal/core/repository"
	"sync"
	"time"
)

type InmemStore struct {
	mu   sync.RWMutex
	data map[string]repository.Port
}

func New() *InmemStore {
	return &InmemStore{
		mu:   sync.RWMutex{},
		data: make(map[string]repository.Port),
	}
}

func (s *InmemStore) CreateOrUpdatePort(ctx context.Context, storePort repository.Port) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[storePort.Id]
	if exists {
		return s.updatePort(ctx, storePort)
	} else {
		return s.createPort(ctx, storePort)
	}
}

func (s *InmemStore) createPort(_ context.Context, port repository.Port) error {
	// set created and updated at
	port.CreatedAt = time.Now()
	port.UpdatedAt = port.CreatedAt

	s.data[port.Id] = port

	return nil
}

func (s *InmemStore) updatePort(_ context.Context, port repository.Port) error {
	// check if port exists
	storePort, exists := s.data[port.Id]
	if !exists {
		return domain.ErrNotFound
	}

	port.CreatedAt = storePort.CreatedAt
	port.UpdatedAt = time.Now()

	s.data[port.Id] = port

	return nil
}

// method for testing purpose
func (s *InmemStore) GetMap() map[string]repository.Port {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}
