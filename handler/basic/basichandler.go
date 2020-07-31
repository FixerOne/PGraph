package basic

import (
	contLocation "pgraph/controller/location"
	repoLocation "pgraph/repository/location"
	servLocation "pgraph/service/location"

	gin "github.com/gin-gonic/gin"
)

var (
	locationRepository repoLocation.Repository         = repoLocation.New()
	locationService    servLocation.Service            = servLocation.New(locationRepository)
	locationController contLocation.LocationController = contLocation.New(locationService)
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
		c.JSON(200, locationController.FindAllCountries())
	})
	h.server.GET("/location/GetStates", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllStates())
	})

	h.server.GET("/location/GetStatesByCountry/:id_Country", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_Country")
		c.JSON(200, locationController.FindStatesByCountry(id))
	})

	h.server.GET("/location/GetCitiesByCountry/:id_Country", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_Country")
		c.JSON(200, locationController.FindCitiesByCountry(id))
	})

	h.server.GET("/location/GetCitiesByState/:id_state", func(c *gin.Context) {
		setUpHeaders(c)
		id := c.Param("id_state")
		c.JSON(200, locationController.FindCitiesByState(id))
	})

	h.server.GET("/location/GetSates", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllCountries())
	})

	h.server.GET("/location/GetCities", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.FindAllCities())
	})

	h.server.OPTIONS("/location/UpdateCountry", func(c *gin.Context) {
		setUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/location/UpdateCountry", func(c *gin.Context) {
		setUpHeaders(c)
		c.JSON(200, locationController.UpdateCountry(c))
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
