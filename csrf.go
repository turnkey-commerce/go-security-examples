package main

import (
	"fmt"
	"net/http"
	"text/template"

	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/turnkey-commerce/go-security-examples/utility"
)

var csrfDatastore = "database/csrf.txt"

// CSRF protection for the post requests.
// In production use https and set the csrf.Secure flag to true!
var CSRF = csrf.Protect([]byte("really#long#key#for#cookie"), csrf.Secure(false))

func getCSRFHome(rw http.ResponseWriter, req *http.Request) {
	balance, err := utility.GetDatastoreString(csrfDatastore)
	utility.Check(err)
	vars := map[string]interface{}{
		"Balance": balance,
	}
	t, _ := template.ParseFiles("templates/csrfHome.gohtml")
	t.Execute(rw, vars)
}

func getCSRFInput1(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/csrfInput1.gohtml")
	t.Execute(rw, nil)
}

func postCSRF1(rw http.ResponseWriter, req *http.Request) {
	csrfInput1 := req.PostFormValue("csrfInput1")
	processWithdrawal(csrfInput1)
	http.Redirect(rw, req, "/", http.StatusSeeOther)
}

func getCSRFInput2(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/csrfInput2.gohtml")
	vars := map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(req),
	}
	t.Execute(rw, vars)
}

func postCSRF2(rw http.ResponseWriter, req *http.Request) {
	csrfInput2 := req.PostFormValue("csrfInput2")
	processWithdrawal(csrfInput2)
	http.Redirect(rw, req, "/", http.StatusSeeOther)
}

func getCSRFInput3(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/csrfInput3.gohtml")
	t.Execute(rw, nil)
}

func processWithdrawal(amount string) {
	var balance, withdrawal int64
	balanceStr, err := utility.GetDatastoreString(csrfDatastore)
	utility.Check(err)
	balance, err = strconv.ParseInt(balanceStr, 10, 64)
	utility.Check(err)
	withdrawal, err = strconv.ParseInt(amount, 10, 64)
	utility.Check(err)
	balance -= withdrawal
	err = utility.SaveDatastoreString(csrfDatastore, strconv.FormatInt(balance, 10))
	utility.Check(err)
}

func main() {
	// Initialize the datastore
	err := utility.SaveDatastoreString(csrfDatastore, "50000")
	utility.Check(err)
	// set up routers and route handlers
	r := mux.NewRouter()
	//Protected router is setup separately as a subrouter
	r2 := mux.NewRouter().PathPrefix("/protect").Subrouter()

	r.HandleFunc("/csrfInput1", getCSRFInput1).Methods("GET")
	r.HandleFunc("/csrfPost1", postCSRF1).Methods("POST")
	r2.HandleFunc("/csrfInput2", getCSRFInput2).Methods("GET")
	r2.HandleFunc("/csrfPost2", postCSRF2).Methods("POST")
	// Show the case when the token isn't included in the request when required.
	r2.HandleFunc("/csrfInput3", getCSRFInput3).Methods("GET")
	r.HandleFunc("/", getCSRFHome).Methods("GET")

	http.Handle("/", r)
	// Wrap protected router with CSRF protection
	http.Handle("/protect/", CSRF(r2))

	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", nil)
}
