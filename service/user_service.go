package service

import (
	"final_project_3/repositories"
	"net/mail"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	rr repositories.UserRepoApi
}

func NewUserService(rr repositories.UserRepoApi) *UserService { //provie service
	return &UserService{rr: rr}
}

type UserServiceApi interface {
	UserRegisterService(c *gin.Context) gin.H
	UserLoginService(c *gin.Context) gin.H
	UpdateUserService(c *gin.Context) gin.H
	DeleteUserService(c *gin.Context) gin.H
}

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s UserService) UserRegisterService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	user, err := s.rr.UserRegister(c)
	validateEmail := Valid(user.Email)
	if err != nil {
		result = gin.H{
			"result": "Failed Create User",
		}
	} else if user.Full_name == "" {
		result = gin.H{
			"error": "Your full_name is required",
		}
	} else if user.Email == "" {
		result = gin.H{
			"error": "Your email is required",
		}
	} else if validateEmail == false {
		result = gin.H{
			"error": "Invalid Email Format",
		}
	} else if user.Password == "" {
		result = gin.H{
			"error": "Your password is required",
		}
	} else if len(user.Password) < 6 {
		result = gin.H{
			"error": "Password Minimal 6 Karakter",
		}
	} else if user.Role == "" {
		result = gin.H{
			"error": "Your role is Required",
		}
	} else if user.Role != "admin" && user.Role != "member" {
		result = gin.H{
			"error": "Role Hanya boleh diisi dengan admin atau member",
		}
	} else {
		result = gin.H{
			"id":         user.ID,
			"full_name":  user.Full_name,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		}
	}
	return result
}

func (s UserService) UserLoginService(c *gin.Context) gin.H {
	var result gin.H

	err, comparePass, token := s.rr.UserLogin(c)

	// Validate Email
	if err != nil {
		result = gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		}
	}
	// Validate Password
	if !comparePass {
		result = gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		}
	}
	// Validate Email & Password Jika Berhasil
	if err == nil && comparePass {
		result = gin.H{
			"token": token,
		}
	}

	return result
}

func (us UserService) UpdateUserService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Pengguna, _, err := us.rr.UpdateUser(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			// "Success":    "Data Has been Updated",
			"id":         Pengguna.ID,
			"Full_name":  Pengguna.Full_name,
			"email":      Pengguna.Email,
			"updated_at": Pengguna.UpdatedAt,
		}
	}
	return result
}

func (us UserService) DeleteUserService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := us.rr.DeleteUser(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your account has been successfully deleted",
		}
	}
	return result
}
