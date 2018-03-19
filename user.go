package main

import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
	"time"
)

func setupUserRoutes(router *gin.Engine) { // Changed *gin.Router to *gin.Context

	//endpoint used to register new users
	router.POST("/user", registerUser)

	//endpoint used to fetch user on login
	router.GET("/user", getUser)

	//endpoint used to update user data
	router.PUT("/user", updateUser)
}

/*
 * register the new user and creates http response containing 
 * the user database id and authentication token.
 */
func registerUser(ctx *gin.Context) {

	//The standard akamu json response that will be sent in the http response
	jsonResponse := AkamuJsonResponse{}
	//SignUpForm is the struct that contains all information a user must provide to create a new account.
	var signUpForm DBHandler.SignUpForm
	//SignInResponse is the struct that contain all the information given back to user after login or singup
	var signInResponse DBHandler.SignInResponse

	//grabs data from the http post request and bind it to the signUpForm struct
    ctx.BindJSON(&signUpForm)

    //insert user into the Akamu sql database
    err := DBHandler.InsertUser(ctx, signUpForm, &signInResponse)

    if err != nil {
    	jsonResponse.Status = "error"
		jsonResponse.Message = "Failed inserting user:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}
	

    //TODO: Generate appropriated token
	signInResponse.Token = DBHandler.AuthToken{"token value", time.Now()}

	jsonResponse.Status = "success"
	jsonResponse.Message = "register new user successful"
	jsonResponse.Value = &signInResponse

    ctx.JSON(http.StatusOK, jsonResponse)
  
}

func getUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}