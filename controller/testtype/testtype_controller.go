package testtype

import (
	"pgraph/entity"
	service "pgraph/service/testtype"

	"github.com/gin-gonic/gin"
)

//CompanyController inerface
type TestTypeController interface {
	FindOne() entity.Teststypes
	FindByID(id string) entity.Teststypes
	FindAll() []entity.Teststypes
	FindByProjectID(id string) []entity.Teststypes
	Save(ctx *gin.Context) entity.Teststypes
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) TestTypeController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Teststypes {
	var data entity.Teststypes
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Teststypes {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Teststypes {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Teststypes {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByProjectID(id string) []entity.Teststypes {
	return c.service.FindByProjectID(id)
}
