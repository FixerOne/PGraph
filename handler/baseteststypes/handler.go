package baseteststypes

import (
	controller "pgraph/controller/baseteststypes"
	repository "pgraph/repository/baseteststypes"
	service "pgraph/service/baseteststypes"

	headers "pgraph/handler/headers"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository          = repository.New()
	Service    service.Service                = service.New(Repository)
	Controller controller.BaseTypesController = controller.New(Service)
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

	h.server.OPTIONS("/baseteststypes/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/baseteststypes/Add", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

	h.server.GET("/baseteststypes/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.GET("/baseteststypes/find/:id", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByID(id))
	})

}
