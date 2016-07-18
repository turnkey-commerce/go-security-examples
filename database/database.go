package database

import (
	"database/sql"
	// Import the sqlite3 package as blank.
	_ "github.com/mattn/go-sqlite3"
)

// Contact is one of the contacts for a particular site.
type Contact struct {
	ContactID    int64
	Name         string
	EmailAddress string
	SmsNumber    string
}

// GetContactsByContactID get contacts using safe method when parameterize is passed as true
func GetContactsByContactID(db *sql.DB, userQuery string, parameterize bool) ([]Contact, error) {
	var contacts = []Contact{}
	var err error
	var rows *sql.Rows
	if parameterize {
		query := `SELECT ContactID, Name, EmailAddress, SmsNumber
			FROM Contacts
			WHERE ContactID = $1`
		rows, err = db.Query(query, userQuery)
	} else {
		query := `SELECT ContactID, Name, EmailAddress, SmsNumber
			FROM Contacts
			WHERE ContactID = ` + userQuery
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var ContactID int64
		var Name string
		var EmailAddress string
		var SmsNumber string
		err = rows.Scan(&ContactID, &Name, &EmailAddress, &SmsNumber)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, Contact{ContactID: ContactID, Name: Name,
			EmailAddress: EmailAddress, SmsNumber: SmsNumber})
	}

	return contacts, nil
}
