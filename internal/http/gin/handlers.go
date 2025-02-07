package gin

import (
	"net/http"
	"strconv"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/gin-gonic/gin"
)

// Gin handlers
func Handlers(r *gin.Engine, service beer.UseCase) *gin.Engine {
	r.POST("/v1/beers", storeBeer(service))
	r.PATCH("/v1/beers/:id", updateBeer(service))
	r.DELETE("/v1/beers/:id", removeBeer(service))
	r.GET("/v1/beers/:id", getBeer(service))
	r.GET("/v1/beers", getAllBeers(service))
	return r
}

// Store beer API handler
func storeBeer(service beer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var beer beer.Beer
		if err := c.BindJSON(&beer); err != nil {
			return
		}
		b, err := service.Store(&beer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Beer stored",
			"data": gin.H{
				"id":    b.ID,
				"name":  b.Name,
				"type":  b.Type.String(),
				"style": b.Style.String(),
			},
		})
	}
}

// Update beer API handler
func updateBeer(service beer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid param",
				"error":   err.Error(),
			})
			return
		}
		b, err := service.Get(beer.ID(parsedId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Beer not found",
			})
			return
		}
		if err := c.BindJSON(&b); err != nil {
			return
		}
		b, err = service.Update(b)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Beer updated",
			"data": gin.H{
				"id":    b.ID,
				"name":  b.Name,
				"type":  b.Type.String(),
				"style": b.Style.String(),
			},
		})
	}
}

// Remove beer API handler
func removeBeer(service beer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid param",
				"error":   err.Error(),
			})
			return
		}
		b, err := service.Get(beer.ID(parsedId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Beer not found",
			})
			return
		}
		err = service.Remove(b.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

// Get beer API handler
func getBeer(service beer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid param",
				"error":   err.Error(),
			})
			return
		}
		b, err := service.Get(beer.ID(parsedId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Beer not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Beer found",
			"data": gin.H{
				"id":    b.ID,
				"name":  b.Name,
				"type":  b.Type.String(),
				"style": b.Style.String(),
			},
		})
	}
}

// Get all beers API handler
func getAllBeers(service beer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		bs, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		data := make([]gin.H, 0, len(bs))
		for _, b := range bs {
			data = append(data, gin.H{
				"id":    b.ID,
				"name":  b.Name,
				"type":  b.Type.String(),
				"style": b.Style.String(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Beers found",
			"data":    data,
		})
	}
}
