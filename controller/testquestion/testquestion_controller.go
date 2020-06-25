package testquestion

import (
	"pgraph/entity"
	service "pgraph/service/testquestion"

	"github.com/gin-gonic/gin"
)

//TestQuestionController inerface
type TestQuestionController interface {
	FindOne() entity.Testquestion
	FindByID(id string) entity.Testquestion
	FindAll() []entity.Testquestion
	FindByTestTypeID(id string) []entity.Testquestion
	Save(ctx *gin.Context) entity.Testquestion
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) TestQuestionController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Testquestion {
	var data entity.Testquestion
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Testquestion {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Testquestion {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Testquestion {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Testquestion {
	return c.service.FindByTestTypeID(id)
}
