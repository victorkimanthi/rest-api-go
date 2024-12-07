package models

import (
	"Rest-API/db"
	"Rest-API/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "insert into users (email, password) values (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = userId
	return err
}

func (user *User) ValidateLoginCredentials() error {
	query := "select id,password from users where email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, user.Password)

	if !passwordIsValid {
		return errors.New("Invalid password")
	}

	return nil
}
