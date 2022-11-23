package controllers

import (
	"final_project_3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct { // implementasi Controller
	csa service.CategoriesServiceApi
}

func NewCategoriesController(csa service.CategoriesServiceApi) *CategoriesController {
	return &CategoriesController{csa: csa}
}

func (cc *CategoriesController) CreateCategoriesControllers(c *gin.Context) {
	res := cc.csa.CreateCategoriesService(c)
	c.JSON(http.StatusOK, res)
}

func (cc *CategoriesController) GetAllCategoriesControllers(c *gin.Context) {
	res := cc.csa.GetAllCategoriesService(c)
	c.JSON(http.StatusOK, res)
}

func (cc *CategoriesController) UpdateCategoriesControllers(c *gin.Context) {
	res := cc.csa.UpdateCategoriesService(c)
	c.JSON(http.StatusOK, res)
}

func (cc *CategoriesController) DeleteCategoriesControllers(c *gin.Context) {
	res := cc.csa.DeleteCategoriesService(c)
	c.JSON(http.StatusOK, res)
}
