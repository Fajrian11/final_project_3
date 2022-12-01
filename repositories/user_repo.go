package repositories

import (
	"final_project_3/models"
	"fmt"
	"net/mail"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

type UserRepoApi interface {
	UserRegister(User models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	GetUserByID(userID uint) (models.User, error)
	UpdateUser(User models.User, userID uint) (models.User, error)
	DeleteUser(userID uint) error
}

var (
	appJSON = "application/json"
)

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (ur *UserRepo) GetUserByID(userID uint) (models.User, error) {
	var user models.User
	err := ur.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) UserRegister(User models.User) (models.User, error) {
	err := ur.db.Create(&User).Error
	fmt.Println(User)
	if err != nil {
		fmt.Println(err.Error())
	}

	return User, nil
}

func (ur *UserRepo) UpdateUser(User models.User, userID uint) (models.User, error) {
	err := ur.db.Model(&User).Where("id = ?", userID).Updates(models.User{
		Email:     User.Email,
		Full_name: User.Full_name,
	}).Error

	return User, err
}

func (ur *UserRepo) DeleteUser(userID uint) error {
	err := ur.db.Exec(`
	DELETE users 
	FROM users 
	WHERE users.id = ?`, userID).Error

	return err
}
