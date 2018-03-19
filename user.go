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

	//database type is int(11)
	const initialUserExperience int = 0
	//the int(10) UN is the avatar id
	const initialUserAvatar uint32 = 1
	//the int(10) UN is the title id
	const initialUserTitle uint32 = 1
	//the tinyint(3) UN (sql boolean) is 0 for unverified users
	const initialVerifiedStatus bool = false

	//SignUpForm is the struct that contains all information a user must provide to create a new account.
	signUpForm := SignUpFormular{}

	//grabs data from the http post request and bind it to the signUpForm struct
    err := ctx.BindJSON(&signUpForm)

    //test for erros binding http request data to the signUpForm
    if err != nil {
		ctx.String(http.StatusBadRequest, "Failed binding payload to signUpForm. " + err.Error())
		return
	}

    //creates an empty user struct
    user := DBHandler.User{TimeRegistered:time.Now(), Username:signUpForm.Username, Password:signUpForm.Password,
    		Email:signUpForm.Email, Semester:signUpForm.Semester, Experience:initialUserExperience, SelectedAvatar:initialUserAvatar,
    		SelectedTitle:initialUserTitle, University:signUpForm.University, Verified:initialVerifiedStatus}

    //insert user into the Akamu sql database and returns the id
    id, err := DBHandler.InsertUser(&user)

    if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed inserting user: " + err.Error())
		return
	}

	//set http response
    ctx.JSON(http.StatusOK, gin.H{"id":id})
}

func getUser (ctx *gin.Context) {

    //creates an empty user struct
	user := DBHandler.User{}

	//TODO: authentification. Perhaps also check if the user id given is the one used to make the token

	//parse the id string into an int id
	id , err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	//test for erros converting the payload to the id
    if err != nil {
		ctx.String(http.StatusBadRequest, "Failed getting the user id value from the http request. " + err.Error())
		return
	}

	//select user in the database
    err = DBHandler.SelectUserById(uint32(id), &user)

    if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed selecting user from DB. " + err.Error())
		return
	}

	//set http response
    ctx.JSON(http.StatusOK, &user)
}