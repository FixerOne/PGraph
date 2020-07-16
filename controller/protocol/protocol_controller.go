package protocol

import (
	"pgraph/entity"
	service "pgraph/service/protocol"

	"github.com/gin-gonic/gin"
)

//ProtocolController inerface
type ProtocolController interface {
	FindOne() entity.Protocol
	FindByID(id string) entity.Protocol
	FindAll() []entity.Protocol
	Save(ctx *gin.Context) entity.Protocol
	Update(ctx *gin.Context) entity.Protocol
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) ProtocolController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Protocol {
	var dataEntity entity.Protocol
	ctx.ShouldBindJSON(&dataEntity)
	return c.service.Save(dataEntity)
}

//Update method
func (c *controller) Update(ctx *gin.Context) entity.Protocol {
	var dataEntity entity.Protocol
	ctx.ShouldBindJSON(&dataEntity)
	return c.service.Update(dataEntity)
}

//FindAll method
func (c *controller) FindAll() []entity.Protocol {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Protocol {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Protocol {
	return c.service.FindByID(id)
}
