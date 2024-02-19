package models

import (
	"errors"
	"time"

	"github.com/communi-tree/twigs-api/app/utils/services"
	"github.com/gin-gonic/gin"
)

type Login struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" form:"username" gorm:"uniqueIndex" binding:"required"`
	Password  string    `json:"-" form:"password"  binding:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l Login) Find(request *gin.Context) (string, User, error) {
	user := User{}
	if err := request.ShouldBind(&l); err == nil {
		result := DB.Where("username = ?", l.Username).Find(&l)
		if result.Error != nil || result.RowsAffected == 0 || !CheckHashPassword(l.Password, request.PostForm("password")) {
			error_message = errors.New("Record Not Found")
		} else {
			DB.InnerJoins("Login").InnerJoins("Subdivision").Where("login_id = ?", l.ID).Find(&user)
			tokenString, error_message = services.CreateToken(l.Username)
			if error_message != nil {
				error_message = errors.New("Invalid Strings")
			}
		}
	} else {
		error_message = err
	}

	return tokenString, user, error_message
}

func (l Login) Where(request *gin.Context) error {
	return error_message
}

func (l *Login) Create(request *gin.Context) error {
	return error_message
}

func (l Login) update(request *gin.Context) {}

func (l Login) delete(request *gin.Context) {}
