//go:build unit

package beer_test

import (
	"testing"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/artursilveiradev/rest-beer/beer/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServiceStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	b := &beer.Beer{
		ID:    beer.ID(1),
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	r.EXPECT().Store(gomock.Any(), gomock.Any()).Return(b, nil)
	s := beer.NewService(r)

	b, err := s.Store(b)

	assert.NoError(t, err)
	assert.Equal(t, b.ID, beer.ID(1))
	assert.Equal(t, b.Name, "Heineken")
	assert.Equal(t, b.Type, beer.TypeLager)
	assert.Equal(t, b.Style, beer.StylePale)
}

func TestServiceGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	b := &beer.Beer{
		ID:    beer.ID(1),
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	r.EXPECT().Get(gomock.Any(), gomock.Any()).Return(b, nil)
	s := beer.NewService(r)

	b, err := s.Get(b.ID)

	assert.NoError(t, err)
	assert.Equal(t, b.ID, beer.ID(1))
	assert.Equal(t, b.Name, "Heineken")
	assert.Equal(t, b.Type, beer.TypeLager)
	assert.Equal(t, b.Style, beer.StylePale)
}
