package config

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connectionString := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	fmt.Println("Users table checked/created successfully!")
	fmt.Println("Database connected successfully!!")
}
