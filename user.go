package main

import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	//"net/http"
)

func setupUserRoutes(router *gin.Engine) { // Changed *gin.Router to *gin.Context

	//endpoint used to register new users
	router.POST("/user", registerUser)

	//endpoint used to fetch user on login
	router.GET("/user", getUser)

	//endpoint used to update user data
	router.PUT("/user", updateUser)
}

func registerUser(ctx *gin.Context) {

	var signUpForm DBHandler.SignUpForm

	//grabs data from sent in the http post request and bind it to the signUpForm
    ctx.BindJSON(&signUpForm)

    //insert user into the Akamu sql database
    DBHandler.InsertUser(ctx, signUpForm)
}

func getUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}