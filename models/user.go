package models

import "example.com/event-booking-rest-api/db"

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
	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil{
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}
