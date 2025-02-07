package beer

import "context"

// Beer entity id
type ID int

// Beer entity type
type BeerType int

// Beer entity style
type BeerStyle int

// Beer entity
type Beer struct {
	ID    ID
	Name  string
	Type  BeerType
	Style BeerStyle
}

// Repository write operations
type Writer interface {
	Store(ctx context.Context, b *Beer) (*Beer, error)
	Update(ctx context.Context, b *Beer) (*Beer, error)
	Remove(ctx context.Context, id ID) error
}

// Repository read operations
type Reader interface {
	Get(ctx context.Context, id ID) (*Beer, error)
	GetAll(ctx context.Context) ([]*Beer, error)
}

// Repository operations
type Repository interface {
	Writer
	Reader
}

// Beer entity use cases
type UseCase interface {
	Store(b *Beer) (*Beer, error)
	Update(b *Beer) (*Beer, error)
	Get(id ID) (*Beer, error)
	GetAll() ([]*Beer, error)
}

// Beer types
const (
	TypeAle   BeerType = 1
	TypeLager BeerType = 2
	TypeMalt  BeerType = 3
	TypeStout BeerType = 4
)

// Returns a string representation of the beer type
func (beerType BeerType) String() string {
	switch beerType {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknown"
}

// Beer styles
const (
	StyleAmber   BeerStyle = 1
	StyleBlonde  BeerStyle = 2
	StyleBrown   BeerStyle = 3
	StyleCream   BeerStyle = 4
	StyleDark    BeerStyle = 5
	StylePale    BeerStyle = 6
	StyleStrong  BeerStyle = 7
	StyleWheat   BeerStyle = 8
	StyleRed     BeerStyle = 9
	StyleIPA     BeerStyle = 10
	StyleLime    BeerStyle = 11
	StylePilsner BeerStyle = 12
	StyleGolden  BeerStyle = 13
	StyleFruit   BeerStyle = 14
	StyleHoney   BeerStyle = 15
)

// Returns a string representation of the beer style
func (beerStyle BeerStyle) String() string {
	switch beerStyle {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "India Pale Ale"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown"
}
