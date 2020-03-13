package main

import (
	"fmt"
	"github.com/fixerone/pgraph/httpd/endpoint"

	//"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Empezando Main General")
	endpoint.InitAll()

	//r := gin.Default()

}
