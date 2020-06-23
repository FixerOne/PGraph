package project

import (
	"pgraph/entity"
	service "pgraph/service/project"

	"github.com/gin-gonic/gin"
)

//ProjectController inerface
type ProjectController interface {
	FindOne() entity.Project
	FindByID(id string) entity.Project
	FindAll() []entity.Project
	Save(ctx *gin.Context) entity.Project
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) ProjectController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Project {
	var dataEntity entity.Project
	ctx.ShouldBindJSON(&dataEntity)
	return c.service.Save(dataEntity)
}

//FindAll method
func (c *controller) FindAll() []entity.Project {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Project {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Project {
	return c.service.FindByID(id)
}
