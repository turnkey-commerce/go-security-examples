package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/turnkey-commerce/go-security-examples/database"
	"github.com/turnkey-commerce/go-security-examples/utility"
)

var db *sql.DB

func getSQLHome(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/sqlHome.gohtml")
	t.Execute(rw, nil)
}

func getSQLInput1(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/sqlInput1.gohtml")
	t.Execute(rw, nil)
}

func postSQL1(rw http.ResponseWriter, req *http.Request) {
	query := req.PostFormValue("query")
	// Call without parameterizing the query (third argument = false)
	contacts, err := database.GetContactsByContactID(db, query, false)
	utility.Check(err)
	t, _ := template.ParseFiles("templates/sqlResults.gohtml")
	vars := map[string]interface{}{
		"Contacts": contacts,
	}
	t.Execute(rw, vars)
}

func getSQLInput2(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/sqlInput2.gohtml")
	t.Execute(rw, nil)
}

func postSQL2(rw http.ResponseWriter, req *http.Request) {
	query := req.PostFormValue("query")
	// Call without parameterizing the query (third argument = false)
	contacts, err := database.GetContactsByContactID(db, query, true)
	utility.Check(err)
	t, _ := template.ParseFiles("templates/sqlResults.gohtml")
	vars := map[string]interface{}{
		"Contacts": contacts,
	}
	t.Execute(rw, vars)
}

func main() {
	// Setup the main db.
	var err error
	db, err = database.InitializeDB("database/sql_injection.db")
	utility.Check(err)

	// set up routers and route handlers
	r := mux.NewRouter()
	r.HandleFunc("/sqlInput1", getSQLInput1).Methods("GET")
	r.HandleFunc("/sqlPost1", postSQL1).Methods("POST")
	r.HandleFunc("/sqlInput2", getSQLInput2).Methods("GET")
	r.HandleFunc("/sqlPost2", postSQL2).Methods("POST")
	r.HandleFunc("/", getSQLHome).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", nil)

}
