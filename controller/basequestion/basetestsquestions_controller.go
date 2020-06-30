package basequestion

import (
	"pgraph/entity"
	service "pgraph/service/basequestion"

	"github.com/gin-gonic/gin"
)

//BaseQuestionController inerface
type BaseQuestionController interface {
	FindOne() entity.Basetestsquestions
	FindByID(id string) entity.Basetestsquestions
	FindAll() []entity.Basetestsquestions
	FindByTestTypeID(id string) []entity.Basetestsquestions
	Save(ctx *gin.Context) entity.Basetestsquestions
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) BaseQuestionController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Basetestsquestions {
	var data entity.Basetestsquestions
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Basetestsquestions {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Basetestsquestions {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Basetestsquestions {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Basetestsquestions {
	return c.service.FindByTestTypeID(id)
}
