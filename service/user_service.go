package service

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"lets_bid/utils"
)

type User struct {
	ID       uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUser(user *User) {
	// bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	conn := utils.ConnectDb()

	// Insert
	_, err = conn.Query("insert into users (name, email, password) values ($1, $2, $3)", user.Name, user.Email, string(hashedPassword))

	if err != nil {
		panic(err)
	}
	// close db when not in use
	defer conn.Close()
}

func DeleteUser(user *User) {
	conn := utils.ConnectDb()
	// Delete
	if _, err := conn.Query("delete from users where users.id $1", user.ID); err != nil {
		panic(err)
	}
	// close db when not in use
	defer conn.Close()
}

func AuthenticateWithPassword(user *User) (uint, error) {
	conn := utils.ConnectDb()

	result := conn.QueryRow("select id, password from users where users.email = $1", user.Email)

	if result != nil {
		return 0, errors.New("User does not exists in database")
	}

	dbResult := &User{}
	err := result.Scan(dbResult.ID, dbResult.Password)

	if err != nil && err == sql.ErrNoRows {
		return 0, errors.New("Wrong credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbResult.Password), []byte(user.Password))

	if err != nil {
		return 0, errors.New("Wrong credentials")
	}
	// close db when not in use
	defer conn.Close()
	return dbResult.ID, nil
}
