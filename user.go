package main

import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
	"time"
	"strconv"
)

func setupUserRoutes(router *gin.Engine) { 

	//endpoint used to register new users
	router.POST("/user", registerUser)
	//endpoint used to fetch the user data
	router.GET("/user/:id", getUser)

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
	signUpForm := DBHandler.SignUpForm{}
	//SignInResponse is the struct that contain all the information given back to user after login or singup
	signInResponse := DBHandler.SignInResponse{}

	//grabs data from the http post request and bind it to the signUpForm struct
    ctx.BindJSON(&signUpForm)

    //insert user into the Akamu sql database
    err := DBHandler.InsertUser(signUpForm, &signInResponse)

    if err != nil {
    	jsonResponse.Status = "error"
		jsonResponse.Message = "Failed inserting user: " + err.Error()
		ctx.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}
	
    //TODO: Generate appropriated token
	signInResponse.Token = DBHandler.AuthToken{"token value", time.Now()}

	//set values in the akamu standard response json
	jsonResponse.Status = "success"
	jsonResponse.Message = "register new user successful"
	jsonResponse.Value = &signInResponse

	//set http response
    ctx.JSON(http.StatusOK, jsonResponse)
}

func getUser (ctx *gin.Context) {

	//The standard akamu json response that will be sent in the http response
	jsonResponse := AkamuJsonResponse{}
	user := DBHandler.User{}

	//TODO: authentification. Perhaps also check if the user id given is the one used to make the token

	//parse the id string into an int id
	id , _ := strconv.ParseInt(ctx.Param("id"), 10, 0)

	//select user in the database
    err := DBHandler.SelectUserById(id, &user)

    if err != nil {
    	jsonResponse.Status = "error"
		jsonResponse.Message = "Failed retrieving user from database: " + err.Error()
		ctx.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	//set values in the akamu standard response json
	jsonResponse.Status = "success"
	jsonResponse.Message = "get user by id successful"
	jsonResponse.Value = &user

	//set http response
    ctx.JSON(http.StatusOK, jsonResponse)
}