package company

import (
	controller "pgraph/controller/company"
	repository "pgraph/repository/company"
	service "pgraph/service/company"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository        = repository.New()
	Service    service.Service              = service.New(Repository)
	Controller controller.CompanyController = controller.New(Service)
)

//Handler interface
type Handler interface {
	StartHandlers()
}

type handler struct {
	server *gin.Engine
}

//New constructor
func New(server *gin.Engine) Handler {
	return &handler{
		server: server,
	}
}

//StartHandlers function
func (h *handler) StartHandlers() {

	h.server.GET("/company/find/:id", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByID(id))
	})

	h.server.GET("/company/GetAll", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.POST("/company/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

	h.server.OPTIONS("/company/Add", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

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
