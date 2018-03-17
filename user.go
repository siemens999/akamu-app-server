package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupUserRoutes(router *gin.Engine) { // Changed *gin.Router to *gin.Context
	router.GET("/user/:user", getUser)
	router.POST("/user", registerUser)
	router.PUT("/user", updateUser)
}

func getUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval

	ctx.String(http.StatusOK, "AKAMU GET USER")
}

func registerUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}