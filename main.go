package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	router := gin.Default()

	router.GET("/", getRoot)

	//split up end points in different files for better structure
	setupUserRoutes(router)
	setupDuelRoutes(router)
	setupFlashcardRoutes(router)
	

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

func getRoot(ctx *gin.Context) {
	ctx.String(http.StatusOK, "AKAMU REST API")
}