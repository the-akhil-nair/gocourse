package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	log.Println("Trying to connect to MariaDB.")

	// This is need to be called one time.
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		//panic("Unable to connect to DB", err)
		return nil, err
	}

	log.Println("Connected to the DB.")

	return db, nil
}
