package gin

import (
	"net/http"

	"github.com/artursilveiradev/rest-beer/beer"
	"github.com/gin-gonic/gin"
)

// Gin handlers
func Handlers(r *gin.Engine, service beer.UseCase) *gin.Engine {
	r.POST("/v1/beer", storeBeer(service))
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
