package DBHandler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	//also import squirrel
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

//implement function to query the database

func selectUser() {
	
}

// multiple possibilities:
// git parameters for selectUser, to select only some specific
// or create different functions like:
func selectAllUser() {

}

func selectUserById() {

}

func selectUserByName() {

}

type SIGNUP struct{
    USER string `json:"user" binding:"required"`
    PASSWORD string `json:"password" binding:"required"`
    EMAIL string `json:"email" binding:"required"`
    UNIVERSITY string `json:"university" binding:"required"`
}
// more functions for other database operations
func InsertUser(ctx *gin.Context, signup SIGNUP) {
	//Test DB Functionality
        db, err := sql.Open("mysql", "root:13abUtv0@/akamu")

        if err != nil {
			ctx.String(http.StatusInternalServerError, "Could not open database connection.")
			return
		}
		if db.Ping() != nil {
			ctx.String(http.StatusInternalServerError, "Could not open database connection.")
			return
		}
		defer db.Close()


		stmt, err := db.Prepare("INSERT INTO  user (time_registered, name, password, email, semester, experience, selected_avatar, selected_title, verified, university, idmongo) VALUES ( 20060102150405, ?, ?, ?, 1, 1, 1, 1, 1, 'test', 'test')")
		
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to prepare query statement.", err)
			return
		}
		defer stmt.Close()

		_ , err = stmt.Exec(signup.USER, signup.PASSWORD, signup.EMAIL)

		ctx.String(http.StatusInternalServerError, "Failed to execute query statement.", err)





}

//etc
