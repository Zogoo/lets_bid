package service

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"lets_bid/utils"
)

type User struct {
	ID       string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUser(user *User) {
	// bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	conn := utils.ConnectDb()

	// Insert
	if _, err = conn.Query("insert into users values ($1, $2)", user.Name, string(hashedPassword)); err != nil {
		panic("Connection error")
	}
}

func DeleteUser(user *User) {
	conn := utils.ConnectDb()
	// Delete
	if _, err := conn.Query("delete from users where users.id $1", user.ID); err != nil {
		panic("Cannot delete users")
	}
}

func AuthenticatePass(user *User) error {
	conn := utils.ConnectDb()

	result := conn.QueryRow("select * from users where users.email = $1", user.Email)

	if result != nil {
		return errors.New("User does not exists in database")
	}

	err := result.Scan(user.Password)

	if err != nil && err == sql.ErrNoRows {
		return errors.New("Wrong credentials")
	}

	return nil
}
