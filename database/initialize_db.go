package database

import (
	"database/sql"
	"fmt"
	"os"
)

// The createStatements is used to initialize the DB with the schema.
// It seems better to put it in the code rather than an external file to
// prevent accidental changes by the users.
const createStatements = `
CREATE TABLE "Contacts" (
	"ContactId"	    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"Name"	        TEXT NOT NULL UNIQUE,
	"EmailAddress"  TEXT NOT NULL,
	"SmsNumber"	    INTEGER
);
INSERT INTO "Contacts" (Name, EmailAddress, SmsNumber)
VALUES ('John Doe', 'JohnDoe@example.com', 5125551212);
INSERT INTO "Contacts" (Name, EmailAddress, SmsNumber)
VALUES ('Jane Doe', 'JaneDoe@example.com', 5125551213);
`

// InitializeDB creates the DB file and the schema if the file doesn't exist.
func InitializeDB(dbPath string) (*sql.DB, error) {
	// Delete the old DB to sart fresh.
	delErr := deleteDb(dbPath)
	if delErr != nil {
		return nil, delErr
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	fmt.Println("New Database, creating Schema...")
	err = createSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// createSchema applies the initial schema creation to the database.
func createSchema(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = db.Exec(createStatements)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// deleteDb removes the DB file.
func deleteDb(dbPath string) error {
	if _, err := os.Stat(dbPath); err == nil {
		err := os.Remove(dbPath)
		if err != nil {
			return err
		}
	}
	return nil
}
