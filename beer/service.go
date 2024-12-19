package beer

import (
	"context"
	"fmt"
)

// Beer entity service
type Service struct {
	repository Repository
}

// Creates a new service
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// Store a beer
func (s *Service) Store(b *Beer) (*Beer, error) {
	ctx := context.Background()
	b, err := s.repository.Store(ctx, b)
	if err != nil {
		return nil, fmt.Errorf("Store repository error: %w", err)
	}
	return b, nil
}

// Get a beer
func (s *Service) Get(id ID) (*Beer, error) {
	ctx := context.Background()
	b, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Get repository error: %w", err)
	}
	return b, nil
}
