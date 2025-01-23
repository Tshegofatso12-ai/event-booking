package models

import (
	"errors"

	"example.com/event-booking/db"
	"example.com/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT into users(email, password) VALUES (?,?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	hahshedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := statement.Exec(u.Email, hahshedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievePassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}
