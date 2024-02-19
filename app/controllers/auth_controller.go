package controllers

import (
	"net/http"

	"github.com/communi-tree/twigs-api/app/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(request *gin.Context) {
	login := models.Login{}
	token, user, e := login.Find(request)
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
