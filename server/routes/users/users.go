package users

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouteV1(ctx context.Context, routeGroup *gin.RouterGroup) {
	userRoutes := routeGroup.Group("/users")
	{
		userRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		userRoutes.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
