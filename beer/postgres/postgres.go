package postgres

import (
	"context"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/jackc/pgx/v5"
)

// Postgres repository
type Postgres struct {
	conn *pgx.Conn
}

// Creates a new Postgres repository
func NewPostgres(c *pgx.Conn) *Postgres {
	return &Postgres{
		conn: c,
	}
}

// Store a beer
func (r *Postgres) Store(ctx context.Context, b *beer.Beer) (*beer.Beer, error) {
	err := r.conn.QueryRow(ctx,
		"INSERT INTO beer (name, type, style) values ($1, $2, $3) RETURNING id, name, type, style",
		b.Name, b.Type, b.Style).Scan(&b.ID, &b.Name, &b.Type, &b.Style)
	return b, err
}
