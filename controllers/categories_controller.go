package controllers

import (
	"final_project_3/helpers"
	"final_project_3/models"
	"final_project_3/service"
	"fmt"
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
	var input models.Categories

	ContentType := helpers.GetContentType(c)
	var err error = nil

	if ContentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Failed Create Category",
		})
		return
	}
	Categories, err := cc.csa.CreateCategoriesService(input)
	if Categories.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your type is required",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success":    "Data Has been created",
			"id":         Categories.ID,
			"title":      Categories.Type,
			"created_at": Categories.CreatedAt,
		})
	}
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
