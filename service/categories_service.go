package service

import (
	"final_project_3/repositories"

	"github.com/gin-gonic/gin"
)

type CategoriesService struct {
	rr repositories.CategoriesRepoApi
}

func NewCategoriesService(rr repositories.CategoriesRepoApi) *CategoriesService { //provie service
	return &CategoriesService{rr: rr}
}

type CategoriesServiceApi interface {
	CreateCategoriesService(c *gin.Context) gin.H
	GetAllCategoriesService(c *gin.Context) gin.H
	GetCategoryByIdService(c *gin.Context) gin.H
	UpdateCategoriesService(c *gin.Context) gin.H
	DeleteCategoriesService(c *gin.Context) gin.H
}

func (cs CategoriesService) CreateCategoriesService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Categories, _ := cs.rr.CreateCategories(c)
	if Categories.Type == "" {
		result = gin.H{
			"error": "Your type is required",
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been created",
			"id":         Categories.ID,
			"title":      Categories.Type,
			"created_at": Categories.CreatedAt,
		}
	}
	return result
}

func (cs CategoriesService) GetAllCategoriesService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllCategories, err := cs.rr.GetAllCategories(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllCategories,
			"count":  len(GetAllCategories),
		}
	}
	return result
}

func (cs CategoriesService) GetCategoryByIdService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetCategory, err := cs.rr.GetCategoryById(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"id":         GetCategory.ID,
			"title":      GetCategory.Type,
			"created_at": GetCategory.CreatedAt,
		}
	}
	return result
}

func (cs CategoriesService) UpdateCategoriesService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Categories, err := cs.rr.UpdateCategories(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else if Categories.Type == "" {
		result = gin.H{
			"error": "Your type is required",
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been Updated",
			"id":         Categories.ID,
			"title":      Categories.Type,
			"updated_at": Categories.UpdatedAt,
		}
	}
	return result
}

func (cs CategoriesService) DeleteCategoriesService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := cs.rr.DeleteCategories(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your Categories has been successfully deleted",
		}
	}
	return result
}
