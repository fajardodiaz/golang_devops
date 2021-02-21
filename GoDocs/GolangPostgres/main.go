package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pablo"
	password = "abc123"
	dbname   = "go_postgres"
)

func main() {
	//Connection String
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//Open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	//close database
	defer db.Close()

	// Insert data #####################################3333
	//Hardcoded
	insertStat := `insert into "students"("name","roll")values('Gerardo',1)`
	_, e := db.Exec(insertStat)
	CheckError(e)
	fmt.Println("Inserted")

	//Dynamic
	insertdbStat := `insert into "students"("name","roll")values($1,$2)`
	_, e := db.Exec(insertdbStat, "Pelotita", 5)
	CheckError(e)
	fmt.Println("Inserted")

	// update ######################################################
	updateStat := `update "students" set "name"=$1, "roll"=$2 where "id"=$3`
	_, e := db.Exec(updateStat, "Estelita", 1, 8)
	CheckError(e)
	fmt.Println("Updated")

	//Delete
	deleteStat := `delete from "students" where id=$1`
	_, e := db.Exec(deleteStat, 9)
	CheckError(e)
	fmt.Println("Deleted")

	//Get data via select
	rows, err := db.Query(`SELECT "id","name" , "roll" FROM "students"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var roll int
		err = rows.Scan(&id, &name, &roll)
		CheckError(err)

		fmt.Printf("ID: %d\nName: %s\nRoll:%d\n\n", id, name, roll)
	}

	CheckError(err)

	//check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected")
}

// CheckError = check the error...
func CheckError(err error) {
	//ADSAsda
	if err != nil {
		panic(err)
	}
}
