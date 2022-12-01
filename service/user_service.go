package service

import (
	"errors"
	"final_project_3/helpers"
	"final_project_3/models"
	"final_project_3/repositories"
	"net/mail"
)

type UserService struct {
	rr repositories.UserRepoApi
}

func NewUserService(rr repositories.UserRepoApi) *UserService { //provie service
	return &UserService{rr: rr}
}

type UserServiceApi interface {
	UserRegisterService(input models.User) (models.User, error)
	UserLoginService(input models.LoginInput) (models.User, error)
	UpdateUserService(userID uint, input models.User) (models.User, error)
	DeleteUserService(userID uint) error
}

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (us UserService) UserRegisterService(input models.User) (models.User, error) {
	user, _ := us.rr.FindUserByEmail(input.Email)
	if user.ID != 0 {
		return user, errors.New("email yang anda daftarkan sudah tersedia")
	}

	user = models.User{}
	user.Email = input.Email
	user.Full_name = input.Full_name
	user.Password = input.Password
	user.Role = "member"

	user, err := us.rr.UserRegister(user)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (us *UserService) UserLoginService(input models.LoginInput) (models.User, error) {
	user, err := us.rr.FindUserByEmail(input.Email)
	if err != nil {
		return user, err
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(input.Password))

	if !comparePass {
		return user, errors.New("Invalid Password")
	}

	return user, nil
}

func (us UserService) UpdateUserService(userID uint, input models.User) (models.User, error) {
	// get user
	user, err := us.rr.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	// Update user
	user, err = us.rr.UpdateUser(input, userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (us UserService) DeleteUserService(userID uint) error {
	// get user
	_, err := us.rr.GetUserByID(userID)
	if err != nil {
		return err
	}
	// delete user
	err = us.rr.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
