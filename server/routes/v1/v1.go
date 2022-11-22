package v1

import (
	"github.com/gin-gonic/gin"

	"gihub.com/olamilekan000/bp-go/server/routes/users"
)

func AddV1Routes(router *gin.Engine) {
	superGroup := router.Group("/api/v1")
	users.UserRouteV1(superGroup)
}
