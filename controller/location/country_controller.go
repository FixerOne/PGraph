package location

import (
	"pgraph/entity"
	service "pgraph/service/location"

	"github.com/gin-gonic/gin"
)

//LocationController inerface
type LocationController interface {
	FindOne() entity.Country
	FindAll() []entity.Country
	Save(ctx *gin.Context) entity.Country
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

//FindAll method
func (c *controller) FindAll() []entity.Country {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Country {
	return c.service.FindAll()[0]
}
