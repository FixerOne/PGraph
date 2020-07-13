package basetestssections

import (
	"pgraph/entity"
	service "pgraph/service/basetestssections"

	"github.com/gin-gonic/gin"
)

//BaseSectionController inerface
type BaseSectionController interface {
	FindOne() entity.Basetestssections
	FindByID(id string) entity.Basetestssections
	FindAll() []entity.Basetestssections
	FindByTestTypeID(id string) []entity.Basetestssections
	Save(ctx *gin.Context) entity.Basetestssections
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) BaseSectionController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Basetestssections {
	var data entity.Basetestssections
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Basetestssections {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Basetestssections {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Basetestssections {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Basetestssections {
	return c.service.FindByTestTypeID(id)
}
