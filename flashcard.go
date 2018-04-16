package main


import (
	"github.com/gin-gonic/gin"
	"./DBHandler"
	"net/http"
)
func setupFlashcardRoutes(router *gin.Engine) {

	//endpoint used to create new flashcards
	router.POST("/flashcard", createCard)

	//endpoint used get a training list of flashcards
	router.GET("/flashcard", getCards)

}

func createCard (ctx *gin.Context) {

	//creates empty FlashCard struct
	card := DBHandler.Flashcard{}

	err := ctx.ShouldBindJSON(&card)

	if err!= nil {
		ctx.String(http.StatusBadRequest, "Failed binding Flashcard json from request. " + err.Error())
		return
	}
	//TODO: put user id from header as the author attribute from the card. Have to find out a more practical way to use it as uint32 than string cast after getting header parameter
	//insert flashcard into the Akamu sql database and returns it's database id
	id, err := DBHandler.InsertFlashcard(&card)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed inserting flashcard: " + err.Error())
		return
	}

	//set http response
	ctx.JSON(http.StatusOK, gin.H{"id":id})
}

/*
 * returns http response with the list of cards corresponding to the given user and subject.
 * for more information check the openapi specification.
 */
func getCards (ctx *gin.Context) {

	list := DBHandler.TrainingList{}

	//use shouldBindQuery instead of BindQuery in order to handle error response yourself.
	err := ctx.ShouldBindQuery(&list)

	if err!= nil {
		ctx.String(http.StatusBadRequest, "Failed binding list query parameters. " + err.Error())
		return
	}

	//select user in the database
	cards, err := DBHandler.SelectFlashCards(list.Author,list.Subject)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed selecting user from DB. " + err.Error())
		return
	}

	//set http response
	ctx.JSON(http.StatusOK, &cards)
}
