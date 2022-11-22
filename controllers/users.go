package controllers

import (
	"gihub.com/olamilekan000/bp-go/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.IEUserService
}

func (uc *UserController) Create(c *gin.Context) {

}

func (uc *UserController) Get(c *gin.Context) {

}

func NewUserController() *UserController {
	return &UserController{}
}
