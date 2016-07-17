package main

import (
	"bufio"
	"bytes"
	"fmt"
	htmlTemplate "html/template"
	"net/http"
	"os"
	textTemplate "text/template"

	"github.com/gorilla/mux"
)

func getXSSHome(rw http.ResponseWriter, req *http.Request) {
	t, _ := htmlTemplate.ParseFiles("templates/xssHome.gohtml")
	t.Execute(rw, nil)
}

func getXSSInput(rw http.ResponseWriter, req *http.Request) {
	t, _ := htmlTemplate.ParseFiles("templates/xssInput.gohtml")
	t.Execute(rw, nil)
}

func postXSS(rw http.ResponseWriter, req *http.Request) {
	xssInput := req.PostFormValue("xssInput")
	f, err := os.Create("database/xss.txt")
	check(err)
	defer f.Close()
	_, err = f.WriteString(xssInput)
	check(err)
	http.Redirect(rw, req, "/", http.StatusSeeOther)
}

func getXSSView1(rw http.ResponseWriter, req *http.Request) {
	xssInput, err := getXSSText("database/xss.txt")
	check(err)
	vars := map[string]interface{}{
		"XSSInput": xssInput,
	}
	t, _ := textTemplate.ParseFiles("templates/xssView.gohtml")
	t.Execute(rw, vars)
}

func getXSSView2(rw http.ResponseWriter, req *http.Request) {
	xssInput, err := getXSSText("database/xss.txt")
	check(err)
	vars := map[string]interface{}{
		"XSSInput": xssInput,
	}
	t, _ := htmlTemplate.ParseFiles("templates/xssView.gohtml")
	t.Execute(rw, vars)
}

func main() {
	// set up routers and route handlers
	r := mux.NewRouter()
	r.HandleFunc("/xssInput", getXSSInput).Methods("GET")
	r.HandleFunc("/xssPost", postXSS).Methods("POST")
	r.HandleFunc("/xssView1", getXSSView1).Methods("GET")
	r.HandleFunc("/xssView2", getXSSView2).Methods("GET")
	r.HandleFunc("/", getXSSHome).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", nil)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getXSSText(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	return buffer.String(), scanner.Err()
}
