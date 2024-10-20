package models

import (
	"errors"
	"eventBooker/db"
	"eventBooker/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid password")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid password")
	}
	return nil
}
