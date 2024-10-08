package main

import (
	"net/http"

	v1 "github.com/communi-tree/twigs-api/app/controllers"
	"github.com/communi-tree/twigs-api/app/utils/middleware"
	"github.com/gin-gonic/gin"
)

func sample(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the twigs api!")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.GET("/sample", sample)
	router.POST("/login", v1.LoginHandler)
	router.POST("/create_user", v1.CreateUser)
	router.POST("/subdivision", v1.CreateSubdivision)

	router.Use(middleware.AuthMiddelware())

	// router.GET("/users", v1.UserIndex)
	router.GET("/user/:id", v1.UserShow)
	return router
}
