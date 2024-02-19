package models

import (
	"errors"
	"time"

	"github.com/communi-tree/twigs-api/app/utils/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID              int         `json:"id" gorm:"primary_key"`
	FirstName       string      `json:"first_name" form:"first_name" binding:"required"`
	MiddleName      string      `json:"middle_name" form:"middle_name" binding:"required"`
	LastName        string      `json:"last_name" form:"last_name" binding:"required"`
	Block           string      `json:"block" form:"block" binding:"required"`
	Lot             string      `json:"lot" form:"lot" binding:"required"`
	Street          string      `json:"street" form:"street" binding:"required"`
	LoginID         int         `json:"-"`
	Login           Login       `json:"login" gorm:"foreignKey:LoginID;references:ID"`
	SubdivisionUUID string      `json:"-"  form:"subdivision_uuid" binding:"required,uuid"`
	Subdivision     Subdivision `json:"subdivision" gorm:"foreignKey:SubdivisionUUID;references:UUID"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func (u *User) Find(request *gin.Context, user_id int) {
	DB.InnerJoins("Login").InnerJoins("Subdivision").First(&u, user_id)
}

func (u *User) Where(request *gin.Context) {}

func (u *User) Create(request *gin.Context) (string, error) {
	var error_message_login error
	var error_message_user error
	login := Login{}
	if err := request.ShouldBind(&login); err != nil {
		error_message_login = err
	}
	if err := request.ShouldBind(&u); err != nil {
		error_message_user = err
	}

	if error_message_user == nil && error_message_login == nil {

		DB.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			hash, _ := HashPassword(request.PostForm("password"))
			login := Login{
				Username: request.PostForm("username"),
				Password: hash,
			}
			if err := tx.Create(&login).Error; err != nil {
				error_message_login = err
				return err
			}

			u.FirstName = request.PostForm("first_name")
			u.MiddleName = request.PostForm("middle_name")
			u.LastName = request.PostForm("last_name")
			u.Block = request.PostForm("block")
			u.Lot = request.PostForm("lot")
			u.Street = request.PostForm("street")
			u.SubdivisionUUID = request.PostForm("subdivision_uuid")
			u.Login = login

			if err := tx.Create(&u).Error; err != nil {
				error_message_user = err
				return err
			}

			tx.Where("uuid = ?", request.PostForm("subdivision_uuid")).Find(&u.Subdivision)
			tokenString, error_message = services.CreateToken(login.Username)
			return nil
		})

	}

	error_message = errors.Join(error_message_user, error_message_login)

	return tokenString, error_message

}

func (u *User) update(request *gin.Context) {}

func (u *User) delete(request *gin.Context) {}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(hash, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
