package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "jensonsu"
	password = "bug898beet201"
	dbname   = "ginperson"
)

var SqlDB *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	SqlDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	// 因為sql.Open()不會一直保持connection
	// By calling db.Ping() we force our code to actually open up a connection to the database which will validate whether or not our connection string was 100% correct
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
