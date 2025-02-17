package models

import (
	"Login/config"
	"Login/utils"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

//docker run --name postgres -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=mydb -p 5432:5432 -d postgres

func (u *Users) CreateUser() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (name, email, password, age) VALUES ($1, $2, $3, $4) RETURNING id`
	err = config.DB.QueryRow(query, u.Name, u.Email, hashedPassword, u.Age).Scan(&u.ID)
	fmt.Println(err)
	return err
}

func AuthenticateUsers(user Users) (*Users, error) {
	email := user.Email
	password := user.Password

	var foundUser Users
	query := `SELECT id, name, email, password, age FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, email).Scan(&foundUser.ID, &foundUser.Name, &foundUser.Email, &foundUser.Password, &foundUser.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found") // Return explicit error
		}
		return nil, err
	}

	// Validate password
	if !utils.CheckHashPassword(password, foundUser.Password) {
		return nil, fmt.Errorf("invalid password") // Return error instead of nil
	}

	return &foundUser, nil // Return pointer to user
}
