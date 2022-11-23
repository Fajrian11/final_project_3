package controllers

import (
	"final_project_3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct { // implementasi Controller
	tsa service.TaskServiceApi
}

func NewTaskController(tsa service.TaskServiceApi) *TaskController {
	return &TaskController{tsa: tsa}
}

func (tc *TaskController) CreateTaskControllers(c *gin.Context) {
	res := tc.tsa.CreateTaskService(c)
	c.JSON(http.StatusOK, res)
}
func (tc *TaskController) GetAllTaskControllers(c *gin.Context) {
	res := tc.tsa.GetAllTaskService(c)
	c.JSON(http.StatusOK, res)
}

func (tc *TaskController) UpdateTaskControllers(c *gin.Context) {
	res := tc.tsa.UpdateTaskService(c)
	c.JSON(http.StatusOK, res)
}

func (tc *TaskController) UpdateStatusTaskControllers(c *gin.Context) {
	res := tc.tsa.UpdateStatusTaskService(c)
	c.JSON(http.StatusOK, res)
}
func (tc *TaskController) UpdateCategoryTaskControllers(c *gin.Context) {
	res := tc.tsa.UpdateCategoryTaskService(c)
	c.JSON(http.StatusOK, res)
}
func (tc *TaskController) DeleteTaskControllers(c *gin.Context) {
	res := tc.tsa.DeleteTaskService(c)
	c.JSON(http.StatusOK, res)
}
