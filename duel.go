package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupDuelRoutes(router *gin.Engine) {
	router.GET("/duel", getDuels)
	router.POST("/duel", registerDuel)
	router.PUT("/duel", updateDuel)
}

func getDuels(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
	ctx.String(http.StatusOK, "AKAMU GET DUEL")
}

func registerDuel(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}

func updateDuel(ctx *gin.Context) {
	//do what has to be done for this endpoint
	//user repository for data retrieval
}