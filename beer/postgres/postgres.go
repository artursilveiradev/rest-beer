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

// Update a beer
func (r *Postgres) Update(ctx context.Context, b *beer.Beer) (*beer.Beer, error) {
	err := r.conn.QueryRow(ctx,
		"UPDATE beer SET name = $1, type = $2, style = $3 WHERE id = $4 RETURNING id, name, type, style",
		b.Name, b.Type, b.Style, b.ID).Scan(&b.ID, &b.Name, &b.Type, &b.Style)
	return b, err
}

// Remove a beer
func (r *Postgres) Remove(ctx context.Context, id beer.ID) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM beer WHERE id = $1", id)
	return err
}

// Get a beer
func (r *Postgres) Get(ctx context.Context, id beer.ID) (*beer.Beer, error) {
	var beer beer.Beer
	err := r.conn.QueryRow(ctx,
		"SELECT id, name, type, style FROM beer WHERE id = $1",
		id).Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
	return &beer, err
}

// Get all beers
func (r *Postgres) GetAll(ctx context.Context) ([]*beer.Beer, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, name, type, style FROM beer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bs []*beer.Beer
	for rows.Next() {
		var beer beer.Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
		if err != nil {
			return nil, err
		}
		bs = append(bs, &beer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bs, err
}
