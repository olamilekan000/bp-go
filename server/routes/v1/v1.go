package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"gihub.com/olamilekan000/bp-go/server/routes/users"
)

func AddV1Routes(ctx context.Context, router *gin.Engine) {
	superGroup := router.Group("/api/v1")
	users.UserRouteV1(ctx, superGroup)
}
