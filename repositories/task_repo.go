package repositories

import (
	"errors"
	"final_project_3/helpers"
	"final_project_3/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepo {
	return TaskRepo{
		db: db,
	}
}

type TaskRepoApi interface {
	CreateTask(c *gin.Context) (models.Task, error, error)
	GetAllTask(c *gin.Context) ([]models.Task, error)
	UpdateTask(c *gin.Context) (models.Task, error, models.Task)
	UpdateStatusTask(c *gin.Context) (models.Task, error, models.Task)
	UpdateCategoryTask(c *gin.Context) (models.Task, error, models.Task)
	DeleteTask(c *gin.Context) (models.Task, error)
}

func (tr *TaskRepo) CreateTask(c *gin.Context) (models.Task, error, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Task := models.Task{}
	Categories := models.Categories{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = UserID

	err := tr.db.Debug().Create(&Task).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	err2 := tr.db.Select("id").First(&Categories, uint(Task.CategoryID)).Error

	if err2 != nil {
		err2 = errors.New("Data Not Found")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "category_id Not Found",
			"message": "category_id doesn't exists",
		})
	}
	return Task, err, err2
}

func (tr *TaskRepo) GetAllTask(c *gin.Context) ([]models.Task, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Task := models.Task{}
	UserID := uint(userData["id"].(float64))

	Task.UserID = UserID

	var GetAllTask = []models.Task{}
	// err := tr.db.Model(&models.Task{}).Find(&GetAllTask).Error
	err := tr.db.Preload("User").Find(&GetAllTask).Error
	fmt.Println(err)
	return GetAllTask, err
}

func (tr *TaskRepo) UpdateTask(c *gin.Context) (models.Task, error, models.Task) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID
	Task.ID = uint(taskId)

	Task2 := models.Task{}
	err2 := tr.db.First(&Task2, uint(taskId)).Error
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	err := tr.db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{
		Title:       Task.Title,
		Description: Task.Description,
	}).Error

	return Task, err, Task2
}

func (tr *TaskRepo) UpdateStatusTask(c *gin.Context) (models.Task, error, models.Task) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID
	Task.ID = uint(taskId)

	Task2 := models.Task{}
	err2 := tr.db.First(&Task2, uint(taskId)).Error
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	err := tr.db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{
		Status: Task.Status,
	}).Error

	return Task, err, Task2
}

func (tr *TaskRepo) UpdateCategoryTask(c *gin.Context) (models.Task, error, models.Task) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID
	Task.ID = uint(taskId)

	Task2 := models.Task{}
	err2 := tr.db.First(&Task2, uint(taskId)).Error
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	err := tr.db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{
		CategoryID: Task.CategoryID,
	}).Error

	return Task, err, Task2
}

func (tr *TaskRepo) DeleteTask(c *gin.Context) (models.Task, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))

	Task.UserID = userID
	Task.ID = uint(taskId)

	err := tr.db.Exec(`
	DELETE tasks 
	FROM tasks 
	WHERE tasks.id = ?`, taskId).Error

	return Task, err
}
