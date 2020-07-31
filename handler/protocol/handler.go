package protocol

import (
	controller "pgraph/controller/protocol"
	repository "pgraph/repository/protocol"
	service "pgraph/service/protocol"

	headers "pgraph/handler/headers"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository         = repository.New()
	Service    service.Service               = service.New(Repository)
	Controller controller.ProtocolController = controller.New(Service)
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

	h.server.OPTIONS("/protocol/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/protocol/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

	h.server.OPTIONS("/protocol/Update", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/protocol/Update", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Update(c))
	})

	h.server.GET("/protocol/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.GET("/protocol/find/:id", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByID(id))
	})

}
