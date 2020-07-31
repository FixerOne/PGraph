package test

import (
	"pgraph/entity"
	service "pgraph/service/test"

	"github.com/gin-gonic/gin"
)

//TestController inerface
type TestController interface {
	FindOne() entity.Test
	FindByID(id string) entity.Test
	FindAll() []entity.Test
	Save(ctx *gin.Context) entity.Test
	Update(ctx *gin.Context) entity.Test
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) TestController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Test {
	var dataEntity entity.Test
	ctx.ShouldBindJSON(&dataEntity)
	return c.service.Save(dataEntity)
}

//Update method
func (c *controller) Update(ctx *gin.Context) entity.Test {
	var dataEntity entity.Test
	ctx.ShouldBindJSON(&dataEntity)
	return c.service.Update(dataEntity)
}

//FindAll method
func (c *controller) FindAll() []entity.Test {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Test {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Test {
	return c.service.FindByID(id)
}
