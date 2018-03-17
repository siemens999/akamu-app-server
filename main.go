package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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