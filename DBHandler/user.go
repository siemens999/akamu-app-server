package DBHandler

/*
 * encapsulates all functions that interact with database user objects.
 */


import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import _ "github.com/go-sql-driver/mysql"
import sq "github.com/Masterminds/squirrel"


/*
 * This struct encapsulate all data required to register a new user.
 */
type SignUpForm struct{
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email string `json:"email" binding:"required"`
    University string `json:"university" binding:"required"`
    Semester int `json:"semester" binding:"required"`
}

/* 
 * adds a new row to the user table. User specific attributes are taken from the signUpFrom 
 * while non user specific data is hardcoded here. e.G. score is set to 0.
 */
func InsertUser(ctx *gin.Context, signUpForm SignUpForm) {
	//Test DB Functionality
    db, err := sql.Open("mysql", "root:13abUtv0@/akamu")

    //check for errors opening database
    if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not open database connection.")
		return
	}
	if db.Ping() != nil {
		ctx.String(http.StatusInternalServerError, "Could not open database connection.")
		return
	}
	defer db.Close()

	//prepare querry
	insertSQL, _, err := sq.
    Insert("user").Columns("time_registered", "name", "password","email", "semester", 
    	"experience", "selected_avatar", "selected_title", "verified", "university").
    //values können hier leider nicht gesetz werden. alles wird als "?" interpretiert...
    Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?").
    ToSql()

    //creates sql statement
	stmt, err := db.Prepare(insertSQL)
	
	//check for problems with the given sql statement
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to prepare query statement.", err)
		return
	}
	defer stmt.Close()

	//execute sql statement
	_ , err = stmt.Exec(time.Now(), signUpForm.Username, signUpForm.Password, 
		signUpForm.Email, signUpForm.Semester,1,0,1,0,signUpForm.University)

	//check for erros while executing sql statement
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to prepare query statement.", err)
		return
	}
}


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