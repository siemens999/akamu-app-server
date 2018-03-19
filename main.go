package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * JsonResponse struct encapsulate the format from all JSON responses
 * Note status and message here are not the same as in the 
 * http header
 */
type AkamuJsonResponse struct{
	/*
	 * "success" or "error"
	 */
    Status string

    /*
     * Description of the task executed or error occured. 
     * e.G. "Get user request was successful" or 
     * "Could not get user. Error: no user with username 'banana' was found"
     */ 
    Message string

    /*
	 * The response struct that has the data the client requested.
	 * if there is a response value, pass a reference to it here.
	 * if there is no response value, do not initialize this field.
	 */
    Value interface{} 
}

func main() {
	router := gin.Default()

	router.GET("/", getRoot)

	setupUserRoutes(router)
	setupDuelRoutes(router)
	//... for each endpoint
	//function is defined in the file for the endpoint

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

func getRoot(ctx *gin.Context) {
	ctx.String(http.StatusOK, "AKAMU REST API")
}