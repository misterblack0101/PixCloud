package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	Emal         string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us UserService) Create(email, password string) (*User, error) {
	// email
	email = strings.ToLower(email)
	// generate hashed password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while hashing password %w", err)
	}
	hashedPassword := string(hashed)
	// Connect to db and insert the user
	row := us.DB.QueryRow(`
	INSERT INTO users(email,password_hash)
	VALUES($1,$2) RETURNING id`,
		email, hashedPassword)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error while creating user %w", err)
	}

	return &User{
		Emal:         email,
		PasswordHash: hashedPassword,
		Id:           id,
	}, nil

}
