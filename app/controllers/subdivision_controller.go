package controllers

import (
	"net/http"

	"github.com/communi-tree/twigs-api/app/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Subdivision struct {
}

func CreateSubdivision(request *gin.Context) {
	subdivision := models.Subdivision{
		UUID:       uuid.New(),
		Name:       request.PostForm("name"),
		Barangay:   request.PostForm("barangay"),
		City:       request.PostForm("city"),
		PostalCode: request.PostForm("postal_code"),
	}

	request.ShouldBind(&subdivision)
	result := models.DB.Create(&subdivision)
	println(result.RowsAffected)

	request.JSON(
		http.StatusOK,
		gin.H{
			"subdivision": subdivision,
		},
	)
}
