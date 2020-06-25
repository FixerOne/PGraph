package question

import (
	"pgraph/entity"
	service "pgraph/service/question"

	"github.com/gin-gonic/gin"
)

//QuestionController inerface
type QuestionController interface {
	FindOne() entity.Questions
	FindByID(id string) entity.Questions
	FindAll() []entity.Questions
	FindByTestTypeID(id string) []entity.Questions
	Save(ctx *gin.Context) entity.Questions
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) QuestionController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Questions {
	var data entity.Questions
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Questions {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Questions {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Questions {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Questions {
	return c.service.FindByTestTypeID(id)
}
