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
	_, err = conn.Query("insert into users values ($1, $2, $3)", user.Name, user.Email, string(hashedPassword))

	if err != nil {
		panic("Connection error")
	}
	// close db when not in use
	defer conn.Close()
}

func DeleteUser(user *User) {
	conn := utils.ConnectDb()
	// Delete
	if _, err := conn.Query("delete from users where users.id $1", user.ID); err != nil {
		panic("Cannot delete users")
	}
	// close db when not in use
	defer conn.Close()
}

func AuthenticatePass(user *User) error {
	conn := utils.ConnectDb()

	result := conn.QueryRow("select password from users where users.email = $1", user.Email)

	if result != nil {
		return errors.New("User does not exists in database")
	}

	dbResult := &User{}
	err := result.Scan(dbResult.Password)

	if err != nil && err == sql.ErrNoRows {
		return errors.New("Wrong credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbResult.Password), []byte(user.Password))

	if err != nil {
		return errors.New("Wrong credentials")
	}
	// close db when not in use
	defer conn.Close()
	return nil
}
