package utils

import "database/sql"

var conn *sql.DB

func connectDb() {
	var err error
	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
	conn, err = sql.Open("postgres", "dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
}
