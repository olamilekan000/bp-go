package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "gihub.com/olamilekan000/bp-go/server/routes/v1"
)

func AddRoutes(ctx context.Context, router *gin.Engine) {
	v1.AddV1Routes(ctx, router)

	router.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
