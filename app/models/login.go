package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Login struct {
	ID        int       `gorm:"primary_key"`
	Username  string    `json:"username" form:"username" binding:"required"`
	Password  string    `json:"password" form:"password"  binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l Login) Find(request *gin.Context) {
	request.ShouldBind(&l)
}

func (l Login) Where(request *gin.Context) {}

func (l *Login) Create(request *gin.Context) {
	l.Username = request.PostForm("username")
	l.Password = request.PostForm("password")
	DB.Create(&l)
}

func (l Login) update(request *gin.Context) {}

func (l Login) delete(request *gin.Context) {}
