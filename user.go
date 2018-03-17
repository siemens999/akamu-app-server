package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
)

type SIGNUP struct{
    USER string `json:"user" binding:"required"`
    PASSWORD string `json:"password" binding:"required"`
    EMAIL string `json:"password" binding:"required"`
    UNIVERSITY string `json:"password" binding:"required"`
}

func setupUserRoutes(router *gin.Engine) { // Changed *gin.Router to *gin.Context
	router.GET("/user", getUser)
	router.POST("/user", registerUser)
	router.PUT("/user", updateUser)
}

func getUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func registerUser(ctx *gin.Context) {
	var signup SIGNUP
        ctx.BindJSON(&signup)
        //how to call func in files in sub directories
        //insertUser()
        ctx.JSON(200, gin.H{"status": signup.USER})
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}