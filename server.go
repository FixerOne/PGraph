package main

import (
	"io"
	"os"

	contBaseQuestion "pgraph/controller/basequestion"
	repoBaseQuestion "pgraph/repository/basequestion"
	servBaseQuestion "pgraph/service/basequestion"

	contDocumentsType "pgraph/controller/documentstype"
	repoDocumentsType "pgraph/repository/documentstype"
	servDocumentsType "pgraph/service/documentstype"

	contBaseSections "pgraph/controller/basetestssections"
	repoBaseSections "pgraph/repository/basetestssections"
	servBaseSections "pgraph/service/basetestssections"

	contBaseTestsTypes "pgraph/controller/baseteststypes"
	repoBaseTestsTypes "pgraph/repository/baseteststypes"
	servBaseTestsTypes "pgraph/service/baseteststypes"

	handlerBasic "pgraph/handler/basic"
	handlerCompany "pgraph/handler/company"
	handlerProject "pgraph/handler/project"
	handlerProtocol "pgraph/handler/protocol"
	handlerTest "pgraph/handler/test"
	handlerTestquestion "pgraph/handler/testquestion"
	handlerTestType "pgraph/handler/testtype"
	handlerUser "pgraph/handler/user"

	"github.com/gin-gonic/gin"
)

var (
	/*companyRepository repoCompany.Repository        = repoCompany.New()
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

	testRepository repoTest.Repository     = repoTest.New()
	testService    servTest.Service        = servTest.New(testRepository)
	testController contTest.TestController = contTest.New(testService)

	protocolRepository repoProtocol.Repository         = repoProtocol.New()
	protocolService    servProtocol.Service            = servProtocol.New(protocolRepository)
	protocolController contProtocol.ProtocolController = contProtocol.New(protocolService)

	testtypeRepository repoTesttype.Repository         = repoTesttype.New()
	testtypeService    servTesttype.Service            = servTesttype.New(testtypeRepository)
	testtypeController contTesttype.TestTypeController = contTesttype.New(testtypeService)

	testquestionRepository repoTestquestion.Repository             = repoTestquestion.New()
	testquestionService    servTestquestion.Service                = servTestquestion.New(testquestionRepository)
	testquestionController contTestquestion.TestQuestionController = contTestquestion.New(testquestionService)*/

	basequestionRepository repoBaseQuestion.Repository             = repoBaseQuestion.New()
	basequestionService    servBaseQuestion.Service                = servBaseQuestion.New(basequestionRepository)
	basequestionController contBaseQuestion.BaseQuestionController = contBaseQuestion.New(basequestionService)

	documentstypeRepository repoDocumentsType.Repository               = repoDocumentsType.New()
	documentstypeService    servDocumentsType.Service                  = servDocumentsType.New(documentstypeRepository)
	documentstypeController contDocumentsType.DocumentsTypesController = contDocumentsType.New(documentstypeService)

	basesectionsRepository repoBaseSections.Repository            = repoBaseSections.New()
	basesectionsService    servBaseSections.Service               = servBaseSections.New(basesectionsRepository)
	basesectionsController contBaseSections.BaseSectionController = contBaseSections.New(basesectionsService)

	baseteststypesRepository repoBaseTestsTypes.Repository          = repoBaseTestsTypes.New()
	baseteststypesService    servBaseTestsTypes.Service             = servBaseTestsTypes.New(baseteststypesRepository)
	baseteststypesController contBaseTestsTypes.BaseTypesController = contBaseTestsTypes.New(baseteststypesService)
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

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()

}

func main() {

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	basicHandler := handlerBasic.New(server)
	basicHandler.StartHandlers()

	companyHandler := handlerCompany.New(server)
	companyHandler.StartHandlers()

	userHandler := handlerUser.New(server)
	userHandler.StartHandlers()

	projectHandler := handlerProject.New(server)
	projectHandler.StartHandlers()

	testHandler := handlerTest.New(server)
	testHandler.StartHandlers()

	protocolHandler := handlerProtocol.New(server)
	protocolHandler.StartHandlers()

	testtypeHandler := handlerTestType.New(server)
	testtypeHandler.StartHandlers()

	testquestionHandler := handlerTestquestion.New(server)
	testquestionHandler.StartHandlers()

	/*server.GET("/company/find/:id", func(c *gin.Context) {
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

	server.OPTIONS("/company/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
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

	server.GET("/test/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, testController.FindAll())
	})

	server.OPTIONS("/protocol/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/protocol/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, protocolController.Save(c))
	})

	server.OPTIONS("/protocol/Update", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/protocol/Update", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, protocolController.Update(c))
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
	})*/

	server.GET("/basequestion/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, basequestionController.FindByID(id))
	})

	server.GET("/basequestion/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, basequestionController.FindAll())
	})

	server.OPTIONS("/basequestion/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/basequestion/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, basequestionController.Save(c))
	})

	server.OPTIONS("/basesection/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/basesection/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, basesectionsController.Save(c))
	})

	server.GET("/basesection/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, basesectionsController.FindAll())
	})

	server.GET("/basesection/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, basesectionsController.FindByID(id))
	})

	server.OPTIONS("/baseteststypes/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/baseteststypes/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, baseteststypesController.Save(c))
	})

	server.GET("/baseteststypes/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, baseteststypesController.FindAll())
	})

	server.GET("/baseteststypes/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, baseteststypesController.FindByID(id))
	})

	server.GET("/documentstype/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, documentstypeController.FindByID(id))
	})

	server.GET("/documentstype/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, documentstypeController.FindAll())
	})

	server.OPTIONS("/documentstype/Update", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	server.POST("/documentstype/Update", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, documentstypeController.Update(c))
	})

	/*server.OPTIONS("/login", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})*/

	/*server.POST("/login", func(c *gin.Context) {
		setUpHeaders(c)
		var data = userController.Login(c)

		if data.ID == 0 {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &entity.Claims{
			Username: data.Mail,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		var jwtKey = []byte("my_secret_key")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		data.Token = tokenString

		c.JSON(200, data)
		//c.String(200, tokenString)

	})*/

	server.Run(":8686")

}
