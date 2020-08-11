package basetestsadmissions

import (
	controller "pgraph/controller/basetestsadmissions"
	repository "pgraph/repository/basetestsadmissions"
	service "pgraph/service/basetestsadmissions"

	headers "pgraph/handler/headers"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository               = repository.New()
	Service    service.Service                     = service.New(Repository)
	Controller controller.BaseAdmissionsController = controller.New(Service)
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

	h.server.GET("/basetestsadmission/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.GET("/basetestsadmission/GetByTestType/:id", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByTestTypeID(id))
	})

}
