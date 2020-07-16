package location

import (
	"pgraph/entity"
	service "pgraph/service/location"

	"github.com/gin-gonic/gin"
)

//LocationController inerface
type LocationController interface {
	FindOne() entity.Country
	FindAllCountries() []entity.Country
	FindStatesByCountry(id string) []entity.State
	FindAllStates() []entity.State
	FindCitiesByState(id string) []entity.City
	FindCitiesByCountry(id string) []entity.City
	FindAllCities() []entity.City
	Save(ctx *gin.Context) entity.Country
	UpdateCountry(ctx *gin.Context) entity.Country
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) LocationController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Country {
	var ent entity.Country
	ctx.ShouldBindJSON(&ent)
	return c.service.Save(ent)
}

func (c *controller) UpdateCountry(ctx *gin.Context) entity.Country {
	var ent entity.Country
	ctx.ShouldBindJSON(&ent)
	return c.service.Save(ent)
}

//FindAll method
func (c *controller) FindAllCountries() []entity.Country {
	return c.service.FindAllCountries()
}

//FindAll method
func (c *controller) FindOne() entity.Country {
	return c.service.FindAllCountries()[0]
}

//FindStatesByCountry method
func (c *controller) FindStatesByCountry(id string) []entity.State {
	return c.service.FindStatesByCountry(id)
}

func (c *controller) FindAllStates() []entity.State {
	return c.service.FindAllStates()
}

func (c *controller) FindCitiesByCountry(id string) []entity.City {
	return c.service.FindCitiesByCountry(id)
}

func (c *controller) FindCitiesByState(id string) []entity.City {
	return c.service.FindCitiesByState(id)
}

func (c *controller) FindAllCities() []entity.City {
	return c.service.FindAllCities()
}
