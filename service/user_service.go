package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func createNewUser(user User) error {
	// bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	// Insert
	if _, err = conn.Query("insert into users values ($1, $2)", user.name, string(hashedPassword)); err != nil {
		return errors.New("Connection error")
	}
}

func deleteUser(user User) error {
	// Delete
	if _, err = conn.Query("delete from users where user.id $1", user.id); err != nil {
		return error.New("Cannot delete users")
	}
}
