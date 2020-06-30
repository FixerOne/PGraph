package main

import (
	"io"
	"os"
	contCompany "pgraph/controller/company"
	repoCompany "pgraph/repository/company"
	servCompany "pgraph/service/company"

	contLocation "pgraph/controller/location"
	repoLocation "pgraph/repository/location"
	servLocation "pgraph/service/location"

	contUser "pgraph/controller/user"
	repoUser "pgraph/repository/user"
	servUser "pgraph/service/user"

	contProject "pgraph/controller/project"
	repoProject "pgraph/repository/project"
	servProject "pgraph/service/project"

	contProtocol "pgraph/controller/protocol"
	repoProtocol "pgraph/repository/protocol"
	servProtocol "pgraph/service/protocol"

	contTesttype "pgraph/controller/testtype"
	repoTesttype "pgraph/repository/testtype"
	servTesttype "pgraph/service/testtype"

	contTestquestion "pgraph/controller/testquestion"
	repoTestquestion "pgraph/repository/testquestion"
	servTestquestion "pgraph/service/testquestion"

	contBaseQuestion "pgraph/controller/basequestion"
	repoBaseQuestion "pgraph/repository/basequestion"
	servBaseQuestion "pgraph/service/basequestion"

	"github.com/gin-gonic/gin"
)

var (
	companyRepository repoCompany.Repository        = repoCompany.New()
	companyService    servCompany.Service           = servCompany.New(companyRepository)
	companyController contCompany.CompanyController = contCompany.New(companyService)

	locationRepository repoLocation.Repository         = repoLocation.New()
	locationService    servLocation.Service            = servLocation.New(locationRepository)
	locationController contLocation.LocationController = contLocation.New(locationService)

	userRepository repoUser.Repository        = repoUser.New()
	userService    servUser.Service           = servUser.New(userRepository)
	userController contUser.CompanyController = contUser.New(userService)

	projectRepository repoProject.Repository        = repoProject.New()
	projectService    servProject.Service           = servProject.New(projectRepository)
	projectController contProject.ProjectController = contProject.New(projectService)

	protocolRepository repoProtocol.Repository         = repoProtocol.New()
	protocolService    servProtocol.Service            = servProtocol.New(protocolRepository)
	protocolController contProtocol.ProtocolController = contProtocol.New(protocolService)

	testtypeRepository repoTesttype.Repository         = repoTesttype.New()
	testtypeService    servTesttype.Service            = servTesttype.New(testtypeRepository)
	testtypeController contTesttype.TestTypeController = contTesttype.New(testtypeService)

	testquestionRepository repoTestquestion.Repository             = repoTestquestion.New()
	testquestionService    servTestquestion.Service                = servTestquestion.New(testquestionRepository)
	testquestionController contTestquestion.TestQuestionController = contTestquestion.New(testquestionService)

	basequestionRepository repoBaseQuestion.Repository             = repoBaseQuestion.New()
	basequestionService    servBaseQuestion.Service                = servBaseQuestion.New(basequestionRepository)
	basequestionController contBaseQuestion.BaseQuestionController = contBaseQuestion.New(basequestionService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setUpHeaders(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}

func main() {

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/company/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, companyController.FindByID(id))
	})

	server.GET("/company/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, companyController.FindAll())
	})

	server.POST("/company/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, companyController.Save(c))
	})

	server.GET("/user/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, userController.FindAll())
	})

	server.GET("/user/GetByCompany/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, userController.FindByCompanyID(id))
	})

	server.GET("/project/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, projectController.FindAll())
	})

	server.GET("/project/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, projectController.FindByID(id))
	})

	server.GET("/project/GetByCompany/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, projectController.FindByCompanyID(id))
	})

	server.GET("/location/GetCountries", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllCountries())
	})

	server.GET("/location/GetStatesByCountry/:id_Country", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_Country")
		c.JSON(200, locationController.FindStatesByCountry(id))
	})

	server.GET("/location/GetCitiesByState/:id_state", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_state")
		c.JSON(200, locationController.FindCitiesByState(id))
	})

	server.GET("/location/GetSates", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllCountries())
	})

	server.GET("/location/GetCities", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllCountries())
	})

	server.GET("/protocol/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, protocolController.FindAll())
	})

	server.GET("/protocol/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, protocolController.FindByID(id))
	})

	server.GET("/testType/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, testtypeController.FindAll())
	})

	server.GET("/testType/GetByProject/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, testtypeController.FindByProjectID(id))
	})

	server.GET("/testQuestion/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, testquestionController.FindAll())
	})

	server.GET("/testQuestion/GetByTestType/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, testquestionController.FindByTestTypeID(id))
	})

	server.GET("/basequestion/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, basequestionController.FindAll())
	})

	server.Run(":8686")

}
