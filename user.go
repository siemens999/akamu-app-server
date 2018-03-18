package main

import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
	"encoding/json"
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

func registerUser(ctx *gin.Context) {

	var signUpForm DBHandler.SignUpForm
	var signInResponse DBHandler.SignInResponse

	//grabs data from the http post request and bind it to the signUpForm
    ctx.BindJSON(&signUpForm)

    //insert user into the Akamu sql database
    err := DBHandler.InsertUser(ctx, signUpForm, &signInResponse)

    if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed inserting user:" + err.Error())
		return
	}
	

    //TODO: Generate appropriated token
	signInResponse.Token = DBHandler.AuthToken{"token value", time.Now()}

	//build json
	//res , err := json.Marshal(&signInResponse)

	//if err != nil {
	//	ctx.String(http.StatusInternalServerError, "Failed to create json response:" + err.Error())
	//	return
	//}

	//ctx.JSON(http.StatusOK, string(res))
    ctx.JSON(http.StatusOK, signInResponse)
  
}

func getUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateUser(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}