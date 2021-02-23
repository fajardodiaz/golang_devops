package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Configure Database
	db, err := sql.Open("mysql", "fajardo:@@ABC123abc123@127.0.0.1:3306/go_mysql?parseTime=true")

	//Make sure to check the error
	err = db.Ping()

	//Insert in database

	//Router
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Title: %s Page: %s", title, page)

	})

	//Methods
	// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//Hostnames & Subdomains
	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	//Schemes
	// r.HandleFunc("/secure", SecureHandler).Schemes("https")
	// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	//Path Prefixes & Subrouters
	// bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/{title}", GetBook)

	//Serving Static Assets
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Server
	http.ListenAndServe(":8055", r)
}

func CreateBook() {
	fmt.Println("METHOD POST")
}

func ReadBook() {
	fmt.Println("METHOD POST")
}

func UpdateBook() {
	fmt.Println("METHOD POST")
}

func DeleteBook() {
	fmt.Println("METHOD POST")
}

/**
You can read GET parameters with r.URL.Query().Get("token") or POST parameters
(fields from an HTML form) with r.FormValue("email").
*/
