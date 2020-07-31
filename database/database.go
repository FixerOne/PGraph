package database

import (
	"fmt"

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

	/*username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")*/

	username := "vsa"
	password := "vsa"
	dbName := "vsa"
	dbHost := "localhost"

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
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
