//go:build unit

package postgres_test

import (
	"context"
	"log"
	"testing"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/artursilveiradev/rest-beer/beer/postgres"
	"github.com/artursilveiradev/rest-beer/testhelpers"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostgresRepositoryTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *postgres.Postgres
	ctx         context.Context
}

func (s *PostgresRepositoryTestSuite) SetupSuite() {
	s.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(s.ctx)
	if err != nil {
		log.Fatal(err)
	}
	s.pgContainer = pgContainer
	conn, err := pgx.Connect(s.ctx, s.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.repository = postgres.NewPostgres(conn)
}

func (s *PostgresRepositoryTestSuite) TearDownSuite() {
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (s *PostgresRepositoryTestSuite) TestStoreBeer() {
	t := s.T()
	b := &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	b, err := s.repository.Store(s.ctx, b)

	assert.NoError(t, err)
	assert.NotZero(t, b.ID)
	assert.Equal(t, b.Name, "Heineken")
	assert.Equal(t, b.Type, beer.TypeLager)
	assert.Equal(t, b.Style, beer.StylePale)
}

func (s *PostgresRepositoryTestSuite) TestUpdateBeer() {
	t := s.T()
	b, _ := s.repository.Store(s.ctx, &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	})

	b.Name = "Budweiser"
	b, err := s.repository.Update(s.ctx, b)

	assert.NoError(t, err)
	assert.NotZero(t, b.ID)
	assert.Equal(t, b.Name, "Budweiser")
	assert.Equal(t, b.Type, beer.TypeLager)
	assert.Equal(t, b.Style, beer.StylePale)
}

func (s *PostgresRepositoryTestSuite) TestGetBeer() {
	t := s.T()
	b, _ := s.repository.Store(s.ctx, &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	})

	b, err := s.repository.Get(s.ctx, b.ID)

	assert.NoError(t, err)
	assert.NotZero(t, b.ID)
	assert.Equal(t, b.Name, "Heineken")
	assert.Equal(t, b.Type, beer.TypeLager)
	assert.Equal(t, b.Style, beer.StylePale)
}

func (s *PostgresRepositoryTestSuite) TestGetAllBeers() {
	t := s.T()
	s.repository.Store(s.ctx, &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	})

	bs, err := s.repository.GetAll(s.ctx)

	assert.NoError(t, err)
	assert.Equal(t, len(bs), 1)
	assert.NotZero(t, bs[0].ID)
	assert.Equal(t, bs[0].Name, "Heineken")
	assert.Equal(t, bs[0].Type, beer.TypeLager)
	assert.Equal(t, bs[0].Style, beer.StylePale)
}

func TestPostgresRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresRepositoryTestSuite))
}
