package entity

import (
	//...
	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	//...
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//Claims add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
