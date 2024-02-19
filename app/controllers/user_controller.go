package controllers

import (
	"net/http"
	"strconv"

	"github.com/communi-tree/twigs-api/app/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(request *gin.Context) {

	user := models.User{}

	token, e := user.Create(request)

	if e != nil {
		request.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		request.JSON(
			http.StatusOK,
			gin.H{
				"token": token,
				"user":  user,
			},
		)
	}
}

func UserShow(request *gin.Context) {
	user_id, _ := strconv.Atoi(request.Param("id"))
	user := models.User{}
	user.Find(request, user_id)
	request.JSON(http.StatusOK, gin.H{"data": user})
}
