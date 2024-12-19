package gin

import (
	"net/http"
	"strconv"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/gin-gonic/gin"
)

// Gin handlers
func Handlers(r *gin.Engine, service beer.UseCase) *gin.Engine {
	r.POST("/v1/beer", storeBeer(service))
	r.GET("/v1/beer/:id", getBeer(service))
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
