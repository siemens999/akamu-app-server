package DBHandler

/*
 * encapsulates all functions that interact with database user objects.
 *
 * http logic should not be implemented here. e.G. functions that require "net/http".
 *
 * it is important to understand how transactions, connection pools and other sql db
 * concepts work to avoid serious efficiency problems. Avoid coding database interaction
 * you do not fully understand. 
 *
 * very good tutorial to sql functionality http://mindbowser.com/golang-go-database-sql/
 */


import (
	"database/sql"
	"time"
	"fmt"
)
import (
	_ "github.com/go-sql-driver/mysql"
	//"net/http"
	//"os/user"
	//"log"
	//"os/user"
)
//import sq "github.com/Masterminds/squirrel"

/*
 * the data definition from an user in the sql user table
 * (Does not include mongodb ids)
 */
type User struct{
	Id uint32
	TimeRegistered time.Time
	Username string
	Password string
	Email string
	Semester int 
    Experience int
    SelectedAvatar uint32
    SelectedTitle uint32
    University string
    Verified bool
} 

/* 
 * adds the new user to the database and returns the
 * user database id
 */
func InsertUser(user *User) (id uint32, err error){
	
	//Test DB Functionality
    db, err := sql.Open("mysql", "root:13abUtv0@/akamu")

    //check for errors opening the database
    if err != nil {
		return 0, fmt.Errorf("Could not open database connection." + err.Error())
	}
	//check for errors connecting to the database
	if db.Ping() != nil {
		return 0, fmt.Errorf("Could not open database connection.")
	}
	defer db.Close()

	//create transaction, following a tutorial that did not check for errors here
	tx, _ := db.Begin()
	//defering a Rollback here sounds strange but is advised at http://go-database-sql.org/prepared.html
	defer tx.Rollback()

    //creates the insert sql statement for the transaction
	stmt, err := tx.Prepare("INSERT INTO user (time_registered, username, password, email, semester, experience, selected_avatar, selected_title, verified, university) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	
	//check for problems with the created sql statement
	if err != nil {
		//Rollback transaction in case of error
		tx.Rollback()
		return 0, fmt.Errorf("Failed to prepare query statement. " + err.Error())
	}
	defer stmt.Close()

	//execute sql statement to insert the new user into the user table
	_ , err = stmt.Exec(user.TimeRegistered, user.Username, user.Password, 
		user.Email, user.Semester, user.Experience, user.SelectedAvatar,
		user.SelectedTitle, user.Verified, user.University)

	//check for errors while executing the insert sql statement
	if err != nil {
		//if an error occurred Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Failed executing insert query statement. " + err.Error())
	}

	//creates the sql statement to get the user id in the same transaction
	stmt, err = tx.Prepare("SELECT iduser FROM user WHERE username = ?")
	
    //check for errors creating the sql statement
    if err != nil {
    	//if an error occured Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Could not create statement to get user id. " + err.Error())
	}

	//execute sql query that returns the id from the new user and save its value to SignInResponse.Id
	err = stmt.QueryRow(user.Username).Scan(&id)
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Could not get the id from the user created, rolling back transaction. " + err.Error())
	}

	//check for erros while executing sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Failed executing select id from new user query statement. " + err.Error())
	}
	//commit successful transaction
	tx.Commit()

	//return without errors
	return id, nil
}

/*
 * selects the user with the given id from the user table and
 * returns the respective user struct.
 */
func SelectUserById(id uint32) (User, error) {

	//Test DB Functionality, "?parseTime=true" allows to read time from database
    db, err := sql.Open("mysql", "root:13abUtv0@/akamu?parseTime=true")

    //check for errors opening the database
    if err != nil {
		return User{}, fmt.Errorf("Could not open database connection." + err.Error())
	}
	//check for errors connecting to the database
	if db.Ping() != nil {
		return User{}, fmt.Errorf("Could not open database connection.")
	}
	defer db.Close()

	//create statement to fetch the selected user from db
	stmt, err := db.Prepare("select iduser, time_registered, username, password, email, semester, experience, selected_avatar, selected_title, verified, university FROM user WHERE iduser = ?")
	if err != nil {
		return User{}, fmt.Errorf("Could not prepare sql statement to retrieve selectedUser from Datase. " + err.Error())
	}

	selectedUser := User{}
	//make sql query and save response to the selected user pointer
	err = stmt.QueryRow(id).Scan(&selectedUser.Id, &selectedUser.TimeRegistered, &selectedUser.Username, &selectedUser.Password, &selectedUser.Email, &selectedUser.Semester, &selectedUser.Experience, &selectedUser.SelectedAvatar, &selectedUser.SelectedTitle, &selectedUser.Verified, &selectedUser.University)
	if err != nil {
		return User{}, fmt.Errorf("Could not retrieve user from Datase. " + err.Error())
	}

	//return with no errors
	return selectedUser, nil
}

/*
 * selects all users from the user table and returns the
 * slice containing all the users
 */
func SelectAllUsers() ([]User, error) {

	//Test DB Functionality, "?parseTime=true" allows to read time from database
	db, err := sql.Open("mysql", "root:13abUtv0@/akamu?parseTime=true")

	//check for errors opening the database
	if err != nil {
		return nil,  fmt.Errorf("Could not open database connection." + err.Error())
	}
	//check for errors connecting to the database
	if db.Ping() != nil {
		return nil, fmt.Errorf("Could not open database connection. ")
	}
	defer db.Close()

	//create statement to fetch user from db
	stmt, err := db.Prepare("select iduser, time_registered, username, password, email, semester, experience, selected_avatar, selected_title, verified, university FROM user")
	if err != nil {
		return nil, fmt.Errorf("Could not prepare sql statement to retrieve user from Datase. " + err.Error())
	}

	//execute sql query
	rows, err := stmt.Query()

	if err != nil {
		return  nil, fmt.Errorf("Could not create statement to get user id. " + err.Error())
	}
	defer rows.Close()

	//creates the list of users that will be returned
	var users []User
	//creates a temporary user struct to store data from each row
	var tempUser User

	//iterates through each row returned by the query
	for rows.Next() {
		err = rows.Scan(&(tempUser.Id), &(tempUser.TimeRegistered), &(tempUser.Username),
			&(tempUser.Password), &(tempUser.Email), &(tempUser.Semester),
			&(tempUser.Experience), &(tempUser.SelectedAvatar), &(tempUser.SelectedTitle),
			&(tempUser.Verified), &(tempUser.University))
		if err != nil {
			return  nil, fmt.Errorf("Could not scan db values into user list. " + err.Error())
		}
		users = append(users, tempUser)
	}

	//return with no errors
	return users, nil
}