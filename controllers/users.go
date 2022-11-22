package controllers

import (
	"github.com/gin-gonic/gin"

	"gihub.com/olamilekan000/bp-go/services"
)

type UserController struct {
	service services.IEUserService
}

func (uc *UserController) Create(c *gin.Context) {

}

func (uc *UserController) Get(c *gin.Context) {

}

func NewUserController() *UserController {
	return &UserController{}
}
