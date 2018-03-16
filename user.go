package main

import (
	"github.com/akamu-app-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupUserRoutes(router *gin.Router) { // !TYPE MOST LIKELY WRONG!
	router.GET("/user", getUsers)
	router.POST("/user", registerUser)
	router.PUT("/user", changeUSer)
}

func getUsers(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

// func registerUser
// etc
