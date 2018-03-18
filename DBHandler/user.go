package DBHandler

/*
 * encapsulates all functions that interact with database user objects.
 * should not use classes such as "net/http". http logic should not be 
 * implemented
 */


import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"
import sq "github.com/Masterminds/squirrel"

/*
 * Authentication token
 */
type AuthToken struct{
    Value string `json:"value" binding:"required"`
    Expiriation time.Time `json:"expiration" binding:"required"`
}

/*
 * This struct encapsulate all data returned for users after they signIn or SignUp.
 */
type SignInResponse struct{
    Id string `json:"id" binding:"required"`
    Token AuthToken `json:"token" binding:"required"`
    Email string `json:"email" binding:"required"`
    University string `json:"university" binding:"required"`
    Semester int `json:"semester" binding:"required"`
}
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
func InsertUser(ctx *gin.Context, signUpForm SignUpForm, SignInResponse *SignInResponse) (error){
	//Test DB Functionality
    db, err := sql.Open("mysql", "root:13abUtv0@/akamu")

    //check for errors opening the database
    if err != nil {
		return fmt.Errorf("Could not open database connection." + err.Error())
	}
	//check for errors connecting to the database
	if db.Ping() != nil {
		return fmt.Errorf("Could not open database connection.")
	}
	defer db.Close()

	//create transaction, following tutorial that did not check for errors here
	tx, _ := db.Begin()
	//defering a Rollback sounds strange but is advised at http://go-database-sql.org/prepared.html
	defer tx.Rollback()

	//prepare querry to insert user
	insertSQL, _, err := sq.
    Insert("user").Columns("time_registered", "username", "password","email", "semester", 
    	"experience", "selected_avatar", "selected_title", "verified", "university").
    Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?").
    ToSql()

    //check for errors creating squirrel sinsert sql statement
    if err != nil {
		return fmt.Errorf("Could not create squirrel insert sql statement." + err.Error())
	}

    //creates the insert sql statement for the transaction
	stmt, err := tx.Prepare(insertSQL)
	
	//check for problems with the created sql statement
	if err != nil {
		//Rollback transaction in case of error
		tx.Rollback()
		return fmt.Errorf("Failed to prepare query statement." + err.Error())
	}
	defer stmt.Close()


	//execute sql statement to insert the new user into the user table
	_ , err = stmt.Exec(time.Now(), signUpForm.Username, signUpForm.Password, 
		signUpForm.Email, signUpForm.Semester,0,1,1,0,signUpForm.University)

	//check for erros while executing the insert sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Failed executing insert query statement." + err.Error())
	}

	//prepare querry to get the id using squirrel
	selectSQL, _, err := sq.Select("iduser").From("user").Where(sq.Eq{"username": "?"}).ToSql()

    //check for errors creating squirrel sql statement
    if err != nil {
    	//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Could not create squirrel statement." + err.Error())
	}

	//execute sql query that returns the id from the new user and save its value to SignInResponse.Id
	err = tx.QueryRow(selectSQL,signUpForm.Username).Scan(&SignInResponse.Id)

	//check for erros while executing sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Failed executing select new user id query statement." + err.Error())
	}
	//commit transaction
	tx.Commit()

	//return without errors
	return nil


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