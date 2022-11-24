package service

import (
	"final_project_3/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskService struct {
	rr repositories.TaskRepoApi
}

func NewTaskService(rr repositories.TaskRepoApi) *TaskService { //provie service
	return &TaskService{rr: rr}
}

type TaskServiceApi interface {
	CreateTaskService(c *gin.Context) gin.H
	GetAllTaskService(c *gin.Context) gin.H
	UpdateTaskService(c *gin.Context) gin.H
	UpdateStatusTaskService(c *gin.Context) gin.H
	UpdateCategoryTaskService(c *gin.Context) gin.H
	DeleteTaskService(c *gin.Context) gin.H
}

func (ts TaskService) CreateTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)
	Task, err := ts.rr.CreateTask(c)
	_, err2 := ts.rr.GetCategoryById(Task.CategoryID)
	status := strconv.FormatBool(Task.Status)
	if err != nil {
		result = gin.H{
			"error": "category_id Not Found",
		}
	} else if Task.Title == "" {
		result = gin.H{
			"error": "Your title is required",
		}
	} else if Task.Description == "" {
		result = gin.H{
			"error": "Your description is required",
		}
	} else if status != "false" && status != "true" {
		result = gin.H{
			"error": "Your status is required (true/false)",
		}
	} else if err2 != nil {
		result = gin.H{
			"error":   "category_id Not Found",
			"message": "category_id doesn't exists",
		}
	} else {
		result = gin.H{
			"Success":     "Data Has been created",
			"id":          Task.ID,
			"title":       Task.Title,
			"status":      Task.Status,
			"description": Task.Description,
			"user_id":     Task.UserID,
			"category_id": Task.CategoryID,
			"created_at":  Task.CreatedAt,
		}
	}
	return result
}

func (ts TaskService) GetAllTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllTask, err := ts.rr.GetAllTask(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllTask,
			"count":  len(GetAllTask),
		}
	}
	return result
}

func (ts TaskService) UpdateTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Task, err, Task2 := ts.rr.UpdateTask(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else if Task.Title == "" {
		result = gin.H{
			"error": "Your title is required",
		}
	} else if Task.Description == "" {
		result = gin.H{
			"error": "Your description is required",
		}
	} else {
		result = gin.H{
			"Success":     "Data Has been Updated",
			"id":          Task2.ID,
			"title":       Task.Title,
			"description": Task.Description,
			"status":      Task2.Status,
			"user_id":     Task2.UserID,
			"category_id": Task2.CategoryID,
			"updated_at":  Task.UpdatedAt,
		}
	}
	return result
}

func (ts TaskService) UpdateStatusTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Task, err, Task2 := ts.rr.UpdateTask(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else if Task.Status != false && Task.Status != true {
		result = gin.H{
			"error": "Your status is required (true/false)",
		}
	} else {
		result = gin.H{
			"Success":     "Task Status Has been Updated",
			"id":          Task2.ID,
			"title":       Task2.Title,
			"description": Task2.Description,
			"status":      Task.Status,
			"user_id":     Task2.UserID,
			"category_id": Task2.CategoryID,
			"updated_at":  Task.UpdatedAt,
		}
	}
	return result
}
func (ts TaskService) UpdateCategoryTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Task, err, Task2 := ts.rr.UpdateTask(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else if Task.CategoryID == 0 {
		result = gin.H{
			"error": "Your category_id is required",
		}
	} else {
		result = gin.H{
			"Success":     "Task Category Has been Updated",
			"id":          Task2.ID,
			"title":       Task2.Title,
			"description": Task2.Description,
			"status":      Task2.Status,
			"user_id":     Task2.UserID,
			"category_id": Task.CategoryID,
			"updated_at":  Task.UpdatedAt,
		}
	}
	return result
}

func (cs TaskService) DeleteTaskService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := cs.rr.DeleteTask(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your Task has been successfully deleted",
		}
	}
	return result
}
