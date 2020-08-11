package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

//Init Database
func Init() {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbSSLMode := os.Getenv("ssl_mode")
	dbHostElephant := os.Getenv("db_host_elephant")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName, dbSSLMode, password) //Build connection string

	if dbHostElephant == "1" {
		dbURI = fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, dbHost, username)
	}

	//fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

//Close conection
func Close() {
	defer db.Close()
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
