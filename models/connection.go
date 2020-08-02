package models

//go get github.com/lib/pq
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// TODO: get env
const (
	USER    = "blackhat"
	PASS    = "123456"
	DBNAME  = "supercoin"
	SSLMODE = "disable"
)

// Connect to the database
func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASS, DBNAME, SSLMODE)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

// TestConnection to the database
func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("Database connected!")
}
