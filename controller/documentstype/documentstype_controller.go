package documentstype

import (
	"pgraph/entity"
	service "pgraph/service/documentstype"

	"github.com/gin-gonic/gin"
)

//DocumentsTypesController inerface
type DocumentsTypesController interface {
	FindOne() entity.Documentstypes
	FindByID(id string) entity.Documentstypes
	FindAll() []entity.Documentstypes
	FindByProjectID(id string) []entity.Documentstypes
	Save(ctx *gin.Context) entity.Documentstypes
	Update(ctx *gin.Context) entity.Documentstypes
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) DocumentsTypesController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Documentstypes {
	var data entity.Documentstypes
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

func (c *controller) Update(ctx *gin.Context) entity.Documentstypes {
	var data entity.Documentstypes
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Documentstypes {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Documentstypes {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Documentstypes {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByProjectID(id string) []entity.Documentstypes {
	return c.service.FindByProjectID(id)
}
