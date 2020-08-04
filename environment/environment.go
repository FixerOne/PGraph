package environment

import (
	"log"

	"github.com/joho/godotenv"
)

//SetVariables Database
func SetVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
