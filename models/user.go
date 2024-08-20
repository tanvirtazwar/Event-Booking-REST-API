package models

import (
	"example.com/event-booking-rest-api/db"
	"example.com/event-booking-rest-api/utils"
)

type USER struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user USER) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil{
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil{
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}
