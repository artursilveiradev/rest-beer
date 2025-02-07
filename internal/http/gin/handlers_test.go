//go:build unit

package gin_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/artursilveiradev/rest-beer/beer/mocks"
	g "github.com/artursilveiradev/rest-beer/internal/http/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStoreBeer(t *testing.T) {
	t.Run("StatusCreated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Store(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}
		beerJson, _ := json.Marshal(beer)
		req, _ := http.NewRequest("POST", "/v1/beers", strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, string(
			"{\"data\":{\"id\":1,\"name\":\"Heineken\",\"style\":\"Pale\",\"type\":\"Lager\"},\"message\":\"Beer stored\",\"status\":201}",
		), w.Body.String())
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Store(gomock.Any()).Return(nil, errors.New("internal server error"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}
		beerJson, _ := json.Marshal(beer)
		req, _ := http.NewRequest("POST", "/v1/beers", strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string("{\"message\":\"Internal Server Error\",\"status\":500}"), w.Body.String())
	})
}

func TestUpdateBeer(t *testing.T) {
	t.Run("StatusOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		s.EXPECT().Update(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Budweiser",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name: "Budweiser",
		}
		beerJson, _ := json.Marshal(beer)
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("PATCH", url, strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(
			"{\"data\":{\"id\":1,\"name\":\"Budweiser\",\"style\":\"Pale\",\"type\":\"Lager\"},\"message\":\"Beer updated\",\"status\":200}",
		), w.Body.String())
	})

	t.Run("StatusBadRequest", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name: "Budweiser",
		}
		beerJson, _ := json.Marshal(beer)
		url := fmt.Sprintf("/v1/beers/%s", "foo")
		req, _ := http.NewRequest("PATCH", url, strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(
			"{\"error\":\"strconv.Atoi: parsing \\\"foo\\\": invalid syntax\",\"message\":\"Invalid param\",\"status\":400}",
		), w.Body.String())
	})

	t.Run("StatusNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(nil, errors.New("not found"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name: "Budweiser",
		}
		beerJson, _ := json.Marshal(beer)
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("PATCH", url, strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string("{\"message\":\"Beer not found\",\"status\":404}"), w.Body.String())
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		s.EXPECT().Update(gomock.Any()).Return(nil, errors.New("internal server error"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		beer := beer.Beer{
			Name: "Budweiser",
		}
		beerJson, _ := json.Marshal(beer)
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("PATCH", url, strings.NewReader(string(beerJson)))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string("{\"message\":\"Internal Server Error\",\"status\":500}"), w.Body.String())
	})
}

func TestRemoveBeer(t *testing.T) {
	t.Run("StatusNoContent", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		s.EXPECT().Remove(gomock.Any()).Return(nil)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("DELETE", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("StatusBadRequest", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "foo")
		req, _ := http.NewRequest("DELETE", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(
			"{\"error\":\"strconv.Atoi: parsing \\\"foo\\\": invalid syntax\",\"message\":\"Invalid param\",\"status\":400}",
		), w.Body.String())
	})

	t.Run("StatusNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(nil, errors.New("not found"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("DELETE", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string("{\"message\":\"Beer not found\",\"status\":404}"), w.Body.String())
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		s.EXPECT().Remove(gomock.Any()).Return(errors.New("internal server error"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("DELETE", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string("{\"message\":\"Internal Server Error\",\"status\":500}"), w.Body.String())
	})
}

func TestGetBeer(t *testing.T) {
	t.Run("StatusOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(&beer.Beer{
			ID:    beer.ID(1),
			Name:  "Heineken",
			Type:  beer.TypeLager,
			Style: beer.StylePale,
		}, nil)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(
			"{\"data\":{\"id\":1,\"name\":\"Heineken\",\"style\":\"Pale\",\"type\":\"Lager\"},\"message\":\"Beer found\",\"status\":200}",
		), w.Body.String())
	})

	t.Run("StatusBadRequest", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "foo")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(
			"{\"error\":\"strconv.Atoi: parsing \\\"foo\\\": invalid syntax\",\"message\":\"Invalid param\",\"status\":400}",
		), w.Body.String())
	})

	t.Run("StatusNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().Get(gomock.Any()).Return(nil, errors.New("not found"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/v1/beers/%s", "1")
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string("{\"message\":\"Beer not found\",\"status\":404}"), w.Body.String())
	})
}

func TestGetAllBeers(t *testing.T) {
	t.Run("StatusOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().GetAll().Return([]*beer.Beer{
			{
				ID:    beer.ID(1),
				Name:  "Heineken",
				Type:  beer.TypeLager,
				Style: beer.StylePale,
			},
		}, nil)
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/beers", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(
			"{\"data\":[{\"id\":1,\"name\":\"Heineken\",\"style\":\"Pale\",\"type\":\"Lager\"}],\"message\":\"Beers found\",\"status\":200}",
		), w.Body.String())
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		s := mocks.NewMockUseCase(ctrl)
		s.EXPECT().GetAll().Return(nil, errors.New("internal server error"))
		router := gin.Default()
		router = g.Handlers(router, s)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/beers", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string("{\"message\":\"Internal Server Error\",\"status\":500}"), w.Body.String())
	})
}
