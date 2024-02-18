package controllers

import (
	"net/http"

	"github.com/communi-tree/twigs-api/app/models"
	"github.com/communi-tree/twigs-api/app/utils/services"
	"github.com/gin-gonic/gin"
)

func LoginHandler(request *gin.Context) {
	login := models.Login{}
	if err := request.ShouldBind(&login); err == nil {
		result := models.DB.Where("username = ?", login.Username).Where("password = ?", login.Password).Find(&login)
		if result.Error != nil || result.RowsAffected == 0 {
			request.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Not Found",
				},
			)
		} else {
			tokenString, err := services.CreateToken(login.Username)
			if err != nil {
				panic("Invalid Strings")
			}
			request.JSON(
				http.StatusOK,
				gin.H{
					"token": tokenString,
					"user":  login,
				},
			)
		}

	} else {
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
