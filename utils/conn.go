package utils

import "database/sql"

func ConnectDb() *sql.DB {
	var err error
	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
	conn, err := sql.Open("postgres", "dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
	// close db when not in use
	defer conn.Close()

	return conn
}
