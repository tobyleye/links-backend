package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	UserID    string
}

func OpenDB() (*sql.DB, error) {
	dbConfig := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		DBName:               "links",
		Net:                  "tcp",
		Addr:                 "localhost:3307",
		AllowNativePasswords: true,
	}

	dbConnStr := dbConfig.FormatDSN()

	db, err := sql.Open("mysql", dbConnStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FindUserByEmail(email string) (*User, error) {
	fmt.Println("finding user with email --", email)
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`select user_id, email, password from users where email = ?`, email)

	var user User
	err = row.Scan(&user.UserID, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	// all other types of errors should be reported
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(email string, password string) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	userId := uuid.NewString()

	_, err = db.Query(`insert into users(user_id, email, password) values(?,?,?)`, userId, email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
