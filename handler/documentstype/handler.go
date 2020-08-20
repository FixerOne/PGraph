package documentstype

import (
	controller "pgraph/controller/documentstype"
	repository "pgraph/repository/documentstype"
	service "pgraph/service/documentstype"

	headers "pgraph/handler/headers"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository               = repository.New()
	Service    service.Service                     = service.New(Repository)
	Controller controller.DocumentsTypesController = controller.New(Service)
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

	h.server.GET("/documentstype/find/:id", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByID(id))
	})

	h.server.GET("/documentstype/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.OPTIONS("/documentstype/Save", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/documentstype/Save", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.Save(c))
	})

}
