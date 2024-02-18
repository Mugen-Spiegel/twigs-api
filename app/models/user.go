package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID              int         `json:"id" gorm:"primary_key"`
	FirstName       string      `json:"first_name"`
	MiddleName      string      `json:"middle_name"`
	LastName        string      `json:"last_name"`
	Block           string      `json:"block"`
	Lot             string      `json:"lot"`
	Street          string      `json:"street"`
	LoginID         int         `json:"login_id"`
	Login           Login       `gorm:"foreignKey:LoginID;references:ID"`
	SubdivisionUUID uuid.UUID   `json:"subdivision_id"`
	Subdivision     Subdivision `gorm:"foreignKey:SubdivisionUUID;references:UUID"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func (u *User) Find(request *gin.Context, user_id int) {
	DB.InnerJoins("Login").InnerJoins("Subdivision").First(&u, user_id)
}

func (u *User) Where(request *gin.Context) {}

func (u *User) Create(request *gin.Context, login Login) {
	u.FirstName = request.PostForm("first_name")
	u.MiddleName = request.PostForm("middle_name")
	u.LastName = request.PostForm("last_name")
	u.Block = request.PostForm("block")
	u.Lot = request.PostForm("lot")
	u.Street = request.PostForm("password")
	u.SubdivisionUUID = uuid.MustParse(request.PostForm("subdivision_uuid"))
	u.Login = login
	result := DB.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) update(request *gin.Context) {}

func (u *User) delete(request *gin.Context) {}
