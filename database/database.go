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

	/*e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}*/

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbSSLMode := os.Getenv("ssl_mode")
	dbHostElephant := os.Getenv("db_host_elephant")

	/*username := "vsa"
	password := "vsa"
	dbName := "vsa"
	dbHost := "localhost"*/

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName, dbSSLMode, password) //Build connection string

	if dbHostElephant == "1" {
		//postgres://gblhjfxo:3HaHmRjkMgxFpYJisHPOV4-Jz_IILX91@drona.db.elephantsql.com:5432/gblhjfxo
		dbURI = fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, dbHost, username)
		//dbURI = "postgres://gblhjfxo:3HaHmRjkMgxFpYJisHPOV4-Jz_IILX91@drona.db.elephantsql.com:5432/gblhjfxo"
	}

	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//defer db.Close()

	//db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

//Close conection
func Close() {

	defer db.Close()
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
