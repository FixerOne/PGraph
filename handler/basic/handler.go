package basic

import (
	controller "pgraph/controller/location"
	repository "pgraph/repository/location"
	service "pgraph/service/location"

	gin "github.com/gin-gonic/gin"
)

var (
	Repository repository.Repository         = repository.New()
	Service    service.Service               = service.New(Repository)
	Controller controller.LocationController = controller.New(Service)
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

	h.server.GET("/location/GetCountries", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.FindAllCountries())
	})
	h.server.GET("/location/GetStates", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.FindAllStates())
	})

	h.server.GET("/location/GetStatesByCountry/:id_Country", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_Country")
		c.JSON(200, Controller.FindStatesByCountry(id))
	})

	h.server.GET("/location/GetCitiesByCountry/:id_Country", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_Country")
		c.JSON(200, Controller.FindCitiesByCountry(id))
	})

	h.server.GET("/location/GetCitiesByState/:id_state", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_state")
		c.JSON(200, Controller.FindCitiesByState(id))
	})

	h.server.GET("/location/GetSates", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.FindAllCountries())
	})

	h.server.GET("/location/GetCities", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.FindAllCities())
	})

	h.server.OPTIONS("/location/UpdateCountry", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/location/UpdateCountry", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, Controller.UpdateCountry(c))
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
