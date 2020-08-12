package company

import (
	controller "pgraph/controller/company"
	repository "pgraph/repository/company"
	service "pgraph/service/company"

	headers "pgraph/handler/headers"

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
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByID(id))
	})

	h.server.GET("/company/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.POST("/company/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

	h.server.OPTIONS("/company/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.OPTIONS("/company/Save", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/company/Save", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

}
