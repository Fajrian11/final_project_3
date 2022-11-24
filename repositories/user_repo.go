package repositories

import (
	"final_project_3/helpers"
	"final_project_3/models"
	"fmt"
	"net/mail"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	UserRegister(c *gin.Context) (models.User, error, error)
	UserLogin(c *gin.Context) (error, bool, string)
	UpdateUser(c *gin.Context) (models.User, models.User, error)
	DeleteUser(c *gin.Context) (models.User, error)
}

var (
	appJSON = "application/json"
)

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (ur *UserRepo) UserRegister(c *gin.Context) (models.User, error, error) {
	ContentType := helpers.GetContentType(c)

	JsonUser := models.User{Role: "member"}

	if ContentType == appJSON {
		c.ShouldBindJSON(&JsonUser)
	} else {
		c.ShouldBind(&JsonUser)
	}
	GetUser := models.User{}
	err2 := ur.db.Select("email").First(&GetUser, JsonUser.Email).Error

	err := ur.db.Create(&JsonUser).Error
	fmt.Println(JsonUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	return JsonUser, nil, err2
}

func (ur *UserRepo) UserLogin(c *gin.Context) (error, bool, string) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	// Validate Email
	err := ur.db.Debug().Where("email = ?", User.Email).Take(&User).Error
	// Validate Password
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	// Validate Email & Password Jika Berhasil
	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	return err, comparePass, token
}

func (ur *UserRepo) UpdateUser(c *gin.Context) (models.User, models.User, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Pengguna := models.User{Role: "member"}
	PenggunaDefault := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Pengguna)
	} else {
		c.ShouldBind(&Pengguna)
	}

	Pengguna.ID = userID // Ambil Id dari Claims JWT

	err := ur.db.Model(&Pengguna).Where("id = ?", userID).Updates(models.User{
		Email:     Pengguna.Email,
		Full_name: Pengguna.Full_name,
	}).Error

	return Pengguna, PenggunaDefault, err
}

func (ur *UserRepo) DeleteUser(c *gin.Context) (models.User, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Pengguna := models.User{}

	userID := uint(userData["id"].(float64))

	Pengguna.ID = userID

	err := ur.db.Exec(`
	DELETE users 
	FROM users 
	WHERE users.id = ?`, userID).Error

	return Pengguna, err
}
