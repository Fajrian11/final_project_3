package router

import (
	"final_project_3/config"
	"final_project_3/controllers"
	"final_project_3/database"
	"final_project_3/middlewares"
	"final_project_3/repositories"
	"final_project_3/service"

	"github.com/gin-gonic/gin"
)

func StartAPP() *gin.Engine {
	cfg := config.LoadConfig()
	db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)

	// USER
	userRepo := repositories.NewUserRepo(db)
	userService := service.NewUserService(&userRepo)
	userController := controllers.NewUserController(userService)
	// CATEGORIES
	categoriesRepo := repositories.NewCategoriesRepo(db)
	categoriesService := service.NewCategoriesService(&categoriesRepo)
	categoriesController := controllers.NewCategoriesController(categoriesService)
	// TASK
	taskRepo := repositories.NewTaskRepo(db)
	taskService := service.NewTaskService(&taskRepo)
	taskController := controllers.NewTaskController(taskService)

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegisterControllers)
		userRouter.POST("/login", userController.UserLoginControllers)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/update-account", userController.UpdateUserControllers)
		userRouter.DELETE("/delete-account", userController.DeleteUserControllers)
	}
	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.Use(middlewares.Authentication())
		categoriesRouter.POST("/", middlewares.RoleAuthorization(), categoriesController.CreateCategoriesControllers)
		categoriesRouter.GET("/", categoriesController.GetAllCategoriesControllers)
		categoriesRouter.GET("/:categoryId", categoriesController.GetCategoryByIdControllers)
		categoriesRouter.PATCH("/:categoryId", middlewares.CategoriesAuthorization(), categoriesController.UpdateCategoriesControllers)
		categoriesRouter.DELETE("/:categoryId", middlewares.CategoriesAuthorization(), categoriesController.DeleteCategoriesControllers)
	}
	taskRouter := router.Group("/tasks")
	{
		taskRouter.Use(middlewares.Authentication())
		taskRouter.POST("/", taskController.CreateTaskControllers)
		taskRouter.GET("/", taskController.GetAllTaskControllers)
		taskRouter.PUT("/:taskId", middlewares.TaskAuthorization(), taskController.UpdateTaskControllers)
		taskRouter.PATCH("update-status/:taskId", middlewares.TaskAuthorization(), taskController.UpdateStatusTaskControllers)
		taskRouter.PATCH("update-category/:taskId", middlewares.TaskAuthorization(), taskController.UpdateCategoryTaskControllers)
		taskRouter.DELETE("/:taskId", middlewares.TaskAuthorization(), taskController.DeleteTaskControllers)
	}

	return router
}
