package repositories

import (
	"final_project_3/helpers"
	"final_project_3/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) CategoriesRepo {
	return CategoriesRepo{
		db: db,
	}
}

type CategoriesRepoApi interface {
	CreateCategories(Categories models.Categories) (models.Categories, error)
	GetAllCategories(c *gin.Context) ([]models.Categories, error)
	UpdateCategories(c *gin.Context) (models.Categories, error)
	DeleteCategories(c *gin.Context) (models.Categories, error)
}

func (cr *CategoriesRepo) CreateCategories(Categories models.Categories) (models.Categories, error) {
	err := cr.db.Debug().Create(&Categories).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Categories, err
}

func (cr *CategoriesRepo) GetAllCategories(c *gin.Context) ([]models.Categories, error) {
	var GetAllCategories = []models.Categories{}
	// err := cr.db.Model(&models.Categories{}).Find(&GetAllCategories).Error
	err := cr.db.Model(&models.Categories{}).Preload("Task").Find(&GetAllCategories).Error

	fmt.Println(err)
	return GetAllCategories, err
}

// func (cr *CategoriesRepo) GetCategoryById(c *gin.Context) ([]models.Categories, error) {
// 	var GetAllCategories = []models.Categories{}
// 	// err := cr.db.Model(&models.Categories{}).Find(&GetAllCategories).Error
// 	err := cr.db.Model(&models.Categories{}).Preload("Task").Find(&GetAllCategories).Error

// 	fmt.Println(err)
// 	return GetAllCategories, err
// }

func (cr *CategoriesRepo) UpdateCategories(c *gin.Context) (models.Categories, error) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Categories := models.Categories{}

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	// userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Categories)
	} else {
		c.ShouldBind(&Categories)
	}

	// Categories.UserID = userID
	Categories.ID = uint(categoryId)

	err := cr.db.Model(&Categories).Where("id = ?", categoryId).Updates(models.Categories{
		Type: Categories.Type,
	}).Error

	return Categories, err
}

func (cr *CategoriesRepo) DeleteCategories(c *gin.Context) (models.Categories, error) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	Categories := models.Categories{}

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	// userID := uint(userData["id"].(float64))

	// Photo.UserID = userID
	Categories.ID = uint(categoryId)

	err := cr.db.Exec(`
	DELETE categories
	FROM categories
	WHERE categories.id = ?`, categoryId).Error

	return Categories, err
}
