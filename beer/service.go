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

// Update a beer
func (s *Service) Update(b *Beer) (*Beer, error) {
	ctx := context.Background()
	b, err := s.repository.Update(ctx, b)
	if err != nil {
		return nil, fmt.Errorf("Update repository error: %w", err)
	}
	return b, nil
}

// Remove a beer
func (s *Service) Remove(id ID) error {
	ctx := context.Background()
	err := s.repository.Remove(ctx, id)
	if err != nil {
		return fmt.Errorf("Remove repository error: %w", err)
	}
	return nil
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

// Get all beers
func (s *Service) GetAll() ([]*Beer, error) {
	ctx := context.Background()
	b, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAll repository error: %w", err)
	}
	return b, nil
}
