package db

import (
	"database/sql"
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
	}
}

func ConnectDB() (*sql.DB, error) {
	// Reads environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Builds the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados:", dbname)
	return db, nil
}