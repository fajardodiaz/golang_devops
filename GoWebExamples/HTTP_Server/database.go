package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Configure Database
	db, err := sql.Open("mysql", "fajardo:@@ABC123abc123@(localhost:3306)/go_mysql?parseTime=true")

	//Make sure to check the error
	// err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//Insert in database
	// username := "adiaz"
	// password := "asd134"
	// createdAt := time.Now()

	// //Insert Statement
	// result, err := db.Exec(`INSERT INTO users(username,password,created_at) VALUES(?,?,?)`, username, password, createdAt)
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(id)
	// fmt.Println("Inserted")

	//A single user
	// var (
	// 	id        int
	// 	username  string
	// 	password  string
	// 	createdAt time.Time
	// )

	// //Get Users
	// query := "SELECT id, username, password, created_at FROM users where id=?"
	// if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID: %d \nUsername: %s \nPassword: %s \nCreated at: %v ", id, username, password, createdAt)

	//Query all users
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", users)

	//DELETE
	// _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Server
	http.ListenAndServe(":8055", nil)
}
