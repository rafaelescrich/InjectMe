package models
//go get github.com/lib/pq
import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pq"
)

const (
	USER = "blackhat"
	PASS = "123456"
	DBNAME = "supercoin"
	SSLMODE ="disable"
)

func Connect() *sql.DB {
	URL := fmt.Sprint("user=%s password=%s dbname=%s sslmode=%s", USER, PASS, DBNAME, SSLMODE)
	db, err := sql.Open("postgres",URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection(){
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("Database connected!")
}
