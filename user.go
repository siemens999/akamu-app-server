package main

import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
	"time"
)

func setupUserRoutes(router *gin.Engine) { 

	//endpoint used to register new users
	router.POST("/user", registerUser)

	//TODO implement further endpoints
}

/*
 * register the new user and creates the standard akamu http response. 
 * In case the request is executed successfully the response value is
 * a json containing the user database id and authentication token.
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