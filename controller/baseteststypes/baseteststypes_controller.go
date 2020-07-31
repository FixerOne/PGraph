package baseteststypes

import (
	"pgraph/entity"
	service "pgraph/service/baseteststypes"

	"github.com/gin-gonic/gin"
)

//BaseSectionController inerface
type BaseTypesController interface {
	FindOne() entity.Baseteststypes
	FindByID(id string) entity.Baseteststypes
	FindAll() []entity.Baseteststypes
	FindByTestTypeID(id string) []entity.Baseteststypes
	Save(ctx *gin.Context) entity.Baseteststypes
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) BaseTypesController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Baseteststypes {
	var data entity.Baseteststypes
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Baseteststypes {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Baseteststypes {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Baseteststypes {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Baseteststypes {
	return c.service.FindByTestTypeID(id)
}
