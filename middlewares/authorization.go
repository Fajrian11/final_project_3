package middlewares

import (
	"final_project_3/config"
	"final_project_3/database"
	"final_project_3/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.LoadConfig()
		db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		penggunaId, err := strconv.Atoi(c.Param("penggunaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid paramater",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Pengguna := models.User{}

		err = db.Select("id").First(&Pengguna, uint(penggunaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if Pengguna.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func RoleAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cfg := config.LoadConfig()
		// db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		// categoryId, err := strconv.Atoi(c.Param("categoryId"))
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 		"error":   "Bad Request",
		// 		"message": "invalid paramater",
		// 	})
		// 	return
		// }

		userData := c.MustGet("userData").(jwt.MapClaims)
		role := userData["role"]
		// Categories := models.Categories{}

		// err = db.Select("user_id").First(&Categories, uint(categoryId)).Error

		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		// 		"error":   "Data Not Found",
		// 		"message": "data doesn't exists",
		// 	})
		// 	return
		// }

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CategoriesAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.LoadConfig()
		db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		categoryId, err := strconv.Atoi(c.Param("categoryId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid paramater",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		role := userData["role"]
		Categories := models.Categories{}

		err = db.Select("id").First(&Categories, uint(categoryId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func TaskAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.LoadConfig()
		db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid paramater",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Task := models.Task{}

		err = db.Select("user_id").First(&Task, uint(taskId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if Task.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
