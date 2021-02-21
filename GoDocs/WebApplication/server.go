package main

import (
	"html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

//Global Variable
var templates *template.Template
var client *redis.Client

//Index Page
func indexPage(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange("comments", 0, 20).Result()
	if err != nil {
		return
	}

	templates.ExecuteTemplate(w, "index.html", comments)
}

func main() {
	//Instantiate Redis
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	templates = template.Must(templates.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8500", nil)
}
