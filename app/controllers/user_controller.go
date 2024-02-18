package controllers

import (
	"net/http"
	"strconv"

	"github.com/communi-tree/twigs-api/app/models"
	"github.com/communi-tree/twigs-api/app/utils/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(request *gin.Context) {
	login := models.Login{}
	login.Create(request)
	user := models.User{}
	user.Create(request, login)
	tokenString, _ := services.CreateToken(login.Username)
	request.JSON(
		http.StatusOK,
		gin.H{
			"token": tokenString,
			"user":  user,
		},
	)
}

// func UserIndex(request *gin.Context) {
// 	user := []models.User{}
// 	result := models.DB.Where("id = ?", user_id).Find(&user)
// 	println(result.RowsAffected)
// 	request.JSON(http.StatusOK, gin.H{"data": user})
// }

func UserShow(request *gin.Context) {
	user_id, _ := strconv.Atoi(request.Param("id"))
	user := models.User{}
	user.Find(request, user_id)
	request.JSON(http.StatusOK, gin.H{"data": user})
}
