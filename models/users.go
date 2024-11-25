package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	Email        string
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

	user := User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	err = row.Scan(&user.Id)
	if err != nil {
		return nil, fmt.Errorf("error while creating user %w", err)
	}

	return &user, nil
}

func (us UserService) Login(email, password string) (*User, error) {
	// email
	email = strings.ToLower(email)
	// generate hashed password
	// hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, fmt.Errorf("error while hashing password %w", err)
	// }
	// hashedPassword := string(hashed)
	// Connect to db and insert the user
	row := us.DB.QueryRow(`
	SELECT id , password_hash FROM users WHERE
	email=$1`,
		email)

	user := User{
		Email: email,
	}
	err := row.Scan(&user.Id, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("error while signing in user %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials %w", err)
	}
	fmt.Println(row.Scan())

	return &user, nil
}
