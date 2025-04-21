package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "1234"
	Dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + Dbname)
	return db, nil
}