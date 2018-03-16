package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", getRoot)

	setupUserRoutes(&router)
	setupDeulRoutes(&router)
	//... for each endpoint
	//function is defined in the file for the endpoint

}

func getRoot(ctx *gin.Context) {
	ctx.String(http.StatusOK, "AKAMU REST API")
}
