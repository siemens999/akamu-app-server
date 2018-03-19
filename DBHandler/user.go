package DBHandler

/*
 * encapsulates all functions that interact with database user objects.
 *
 * http logic should not be implemented here. e.G. functions that require "net/http".
 *
 * it is important to understand how transactions, connection pools and other sql db
 * concepts work to avoid serious efficience problems. Avoid coding database interaction 
 * you do not fully understand. 
 *
 * very good tutorial to sql functionality http://mindbowser.com/golang-go-database-sql/
 */


import (
	"database/sql"
	"time"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"
import sq "github.com/Masterminds/squirrel"

//database type is int(11)
const initialUserScore int = 0
//the int(10) UN is the avatar id
const initialUserAvatar uint = 1
//the int(10) UN is the title id
const initialUserTitle uint = 1
//the tinyint(3) UN is 0 for unverified users
const initialVerifiedStatus uint = 0
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
	//the user id
    Id string
    //the user new authentication token
    Token AuthToken
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
 * This struct encapsulate the data from an AppUser
 * TODO: Define final struct. At the moment we have differences
 * between the sql database and the yaml specification
 */
type User struct{
	Id uint
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email string `json:"email" binding:"required"`
    University string `json:"university" binding:"required"`
    Semester int `json:"semester" binding:"required"`
    TimeRegistered time.Time `json:"semester" binding:"required"`
    Experience int `json:"semester" binding:"required"`
    SelectedAvatar int `json:"semester" binding:"required"`
    Verified int `json:"semester" binding:"required"`
    IdMongo string `json:"semester" binding:"required"`
}

/* 
 * adds the new user to the database and saves the respective id value to the signInResponse. 
 * User specific attributes are taken from the signUpForm while non user
 * specific data is hardcoded here. e.G. score is set to 0.
 */
func InsertUser(signUpForm SignUpForm, signInResponse *SignInResponse) (error){
	
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

	//create transaction, following a tutorial that did not check for errors here
	tx, _ := db.Begin()
	//defering a Rollback here sounds strange but is advised at http://go-database-sql.org/prepared.html
	defer tx.Rollback()

	//prepare querry to insert user
	insertSQL, _, err := sq.
    Insert("user").Columns("time_registered", "username", "password","email", "semester", 
    	"experience", "selected_avatar", "selected_title", "verified", "university").
    Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?").
    ToSql()

    //check for errors creating squirrel sinsert sql statement
    if err != nil {
		return fmt.Errorf("Could not create squirrel insert sql statement. " + err.Error())
	}

    //creates the insert sql statement for the transaction
	stmt, err := tx.Prepare(insertSQL)
	
	//check for problems with the created sql statement
	if err != nil {
		//Rollback transaction in case of error
		tx.Rollback()
		return fmt.Errorf("Failed to prepare query statement. " + err.Error())
	}
	defer stmt.Close()

	//execute sql statement to insert the new user into the user table
	_ , err = stmt.Exec(time.Now(), signUpForm.Username, signUpForm.Password, 
		signUpForm.Email, signUpForm.Semester, initialUserScore, initialUserAvatar,
		initialUserTitle, initialVerifiedStatus, signUpForm.University)

	//check for erros while executing the insert sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Failed executing insert query statement. " + err.Error())
	}

	//prepare querry to get the id using squirrel
	selectSQL, _, err := sq.Select("iduser").From("user").Where(sq.Eq{"username": "?"}).ToSql()

    //check for errors creating squirrel sql statement
    if err != nil {
    	//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Could not create squirrel statement. " + err.Error())
	}

	//execute sql query that returns the id from the new user and save its value to SignInResponse.Id
	err = tx.QueryRow(selectSQL,signUpForm.Username).Scan(&signInResponse.Id)

	//check for erros while executing sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return fmt.Errorf("Failed executing select new user id query statement. " + err.Error())
	}
	//commit successful transaction
	tx.Commit()

	//return without errors
	return nil
}

func SelectUserById(id int64, user *User) (error) {

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

	//create statment to fetch user from db
	stmt, err := db.Prepare("select * from user where iduser = ?")
	if err != nil {
		return fmt.Errorf("Could not prepare sql statement to retrieve user from Datase. " + err.Error())
	}

	//make sql query and save response to the user pointer
	err = stmt.QueryRow(id).Scan(user)
	if err != nil {
		return fmt.Errorf("Could not retrieve user from Datase. " + err.Error())
	}
	//return with no errors
	return nil
}