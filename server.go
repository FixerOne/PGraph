package main

import (
	"io"
	"os"
	controller "pgraph/controller/company"
	repository "pgraph/repository/company"
	service "pgraph/service/company"

	"github.com/gin-gonic/gin"
)

var (
	companyRepository repository.Repository        = repository.New()
	companyService    service.CompanyService       = service.New(companyRepository)
	companyController controller.CompanyController = controller.New(companyService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	//setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/GetAll", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(200, companyController.FindAll())
	})

	server.POST("/Add", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(200, companyController.Save(c))
	})

	server.Run(":8686")

}
