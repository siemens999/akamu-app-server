package main


import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
)
func setupFlashcardRoutes(router *gin.Engine) {

	//endpoint used to register new users
	router.POST("/flashcard", createCard)

}

func createCard (ctx *gin.Context) {

	//creates empty FlashCard struct
	card := DBHandler.Flashcard{}

	//grabs data from the http post request and bind it to the FlashCard struct
	err := ctx.BindJSON(&card)

	//test for errors binding http request data to the FlashCard struct
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed binding payload to the FlashCard struct. " + err.Error())
		return
	}

	//insert flashcard into the Akamu sql database and returns the id
	id, err := DBHandler.InsertFlashcard(&card)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed inserting flashcard: " + err.Error())
		return
	}

	//set http response
	ctx.JSON(http.StatusOK, gin.H{"id":id})
}
