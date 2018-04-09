package DBHandler


/*
 * encapsulates all functions that interact with flashcard objects.
 *
 * http logic should not be implemented here. e.G. functions that require "net/http".
 *
 * it is important to understand how transactions, connection pools and other sql db
 * concepts work to avoid serious efficiency problems. Avoid coding database interaction
 * you do not fully understand.
 *
 * good tutorial to sql functionality http://mindbowser.com/golang-go-database-sql/
 */

import (
	//"database/sql"
	//"time"
	//"fmt"
	//"time"
	"time"
	"fmt"
	//"os/user"
	"database/sql"
)

/*
 * the data definition of a flashcard in the sql flashcard table
 */
type Flashcard struct{
	Id uint32
	Author uint32
	Subject uint32
	CreationDate time.Time
	LastModified time.Time
	Version uint32
	FrontText string
	BackText string
	FrontImage uint32
	BackImage uint32
}

/*
 * the data definition of a flashcard training list in the sql traininglist table
 */
type TrainingList struct{
	Author uint32
	Subject uint32
	CreationDate time.Time
	LastModified time.Time
	Version uint32
	UpVotes uint32
	DownVotes uint32
}

func InsertFlashcard(card *Flashcard) (id uint32, err error){
	//Test DB Functionality
	db, err := sql.Open("mysql", "root:13abUtv0@/akamu")

	//check for errors opening the database
	if err != nil {
		return 0, fmt.Errorf("Could not open database connection." + err.Error())
	}
	//check for errors connecting to the database
	if db.Ping() != nil {
		return 0, fmt.Errorf("Could not open database connection. ")
	}
	defer db.Close()

	//create transaction, following a tutorial that did not check for errors here
	tx, _ := db.Begin()
	//deferring a Rollback here sounds strange but is advised at http://go-database-sql.org/prepared.html
	defer tx.Rollback()

	//creates the insert sql statement for the transaction
	stmt, err := tx.Prepare("INSERT INTO flashcard (idflashcard, author, subject, creationdate, lastmodified, version, fronttext, backtext, frontimage, backimage) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	//check for problems with the created sql statement
	if err != nil {
		//Rollback transaction in case of error
		tx.Rollback()
		return 0, fmt.Errorf("Failed to prepare query statement. " + err.Error())
	}
	defer stmt.Close()

	//execute sql statement to insert the new flashcard into the flashcard table
	_ , err = stmt.Exec(card.Id, card.Author, card.Subject, card.CreationDate, card.LastModified, card.Version, card.FrontText,card.BackText, card.FrontImage, card.BackImage)

	//check for errors while executing the insert sql statement
	if err != nil {
		//if an error occurred Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Failed executing insert query statement. " + err.Error())
	}

	//creates the sql statement to get the flashcard id in the same transaction
	stmt, err = tx.Prepare("SELECT LAST_INSERT_ID()")

	//check for errors creating the sql statement
	if err != nil {
		//if an error occured Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Could not create statement to get flashcard id. " + err.Error())
	}

	//execute sql query that returns the id from the new flashcard and save its value to "id" response parameter	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		//if an error occurred, Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Could not get the id from the flashcard created, rolling back the transaction. " + err.Error())
	}

	//check for errors while executing sql statement
	if err != nil {
		//if an error occurred Rollback the transaction
		tx.Rollback()
		return 0, fmt.Errorf("Failed executing select id from new flashcard query statement. " + err.Error())
	}
	//commit successful transaction
	tx.Commit()

	//return without errors
	return id, nil
}