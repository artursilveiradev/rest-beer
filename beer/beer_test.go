//go:build unit

package beer_test

import (
	"testing"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/stretchr/testify/assert"
)

func TestBeerTypeString(t *testing.T) {
	type test struct {
		beerType beer.BeerType
		want     string
	}

	tests := []test{
		{
			beerType: beer.TypeAle,
			want:     "Ale",
		},
		{
			beerType: beer.TypeLager,
			want:     "Lager",
		},
		{
			beerType: beer.TypeMalt,
			want:     "Malt",
		},
		{
			beerType: beer.TypeStout,
			want:     "Stout",
		},
		{
			beerType: 0,
			want:     "Unknown",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.beerType.String(), test.want)
	}
}

func TestBeerStyleString(t *testing.T) {
	type test struct {
		beerStyle beer.BeerStyle
		want      string
	}

	tests := []test{
		{
			beerStyle: beer.StyleAmber,
			want:      "Amber",
		},
		{
			beerStyle: beer.StyleBlonde,
			want:      "Blonde",
		},
		{
			beerStyle: beer.StyleBrown,
			want:      "Brown",
		},
		{
			beerStyle: beer.StyleCream,
			want:      "Cream",
		},
		{
			beerStyle: beer.StyleDark,
			want:      "Dark",
		},
		{
			beerStyle: beer.StylePale,
			want:      "Pale",
		},
		{
			beerStyle: beer.StyleStrong,
			want:      "Strong",
		},
		{
			beerStyle: beer.StyleWheat,
			want:      "Wheat",
		},
		{
			beerStyle: beer.StyleRed,
			want:      "Red",
		},
		{
			beerStyle: beer.StyleIPA,
			want:      "India Pale Ale",
		},
		{
			beerStyle: beer.StyleLime,
			want:      "Lime",
		},
		{
			beerStyle: beer.StylePilsner,
			want:      "Pilsner",
		},
		{
			beerStyle: beer.StyleGolden,
			want:      "Golden",
		},
		{
			beerStyle: beer.StyleFruit,
			want:      "Fruit",
		},
		{
			beerStyle: beer.StyleHoney,
			want:      "Honey",
		},
		{
			beerStyle: 0,
			want:      "Unknown",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.beerStyle.String(), test.want)
	}
}
