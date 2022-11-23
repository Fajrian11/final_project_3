package controllers

import (
	"final_project_3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct { // implementasi Controller
	usa service.UserServiceApi
}

func NewUserController(usa service.UserServiceApi) *UserController {
	return &UserController{usa: usa}
}

func (uc *UserController) UserRegisterControllers(c *gin.Context) {
	res := uc.usa.UserRegisterService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UserLoginControllers(c *gin.Context) {
	res := uc.usa.UserLoginService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateUserControllers(c *gin.Context) {
	res := uc.usa.UpdateUserService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) DeleteUserControllers(c *gin.Context) {
	res := uc.usa.DeleteUserService(c)
	c.JSON(http.StatusOK, res)
}
