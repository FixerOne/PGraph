package basetestsadmissions

import (
	"pgraph/entity"
	service "pgraph/service/basetestsadmissions"

	"github.com/gin-gonic/gin"
)

//BaseAdmissionsController inerface
type BaseAdmissionsController interface {
	FindOne() entity.Basetestsadmissions
	FindByID(id string) entity.Basetestsadmissions
	FindAll() []entity.Basetestsadmissions
	FindByTestTypeID(id string) []entity.Basetestsadmissions
	Save(ctx *gin.Context) entity.Basetestsadmissions
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) BaseAdmissionsController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Basetestsadmissions {
	var data entity.Basetestsadmissions
	ctx.ShouldBindJSON(&data)
	return c.service.Save(data)
}

//FindAll method
func (c *controller) FindAll() []entity.Basetestsadmissions {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindOne() entity.Basetestsadmissions {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Basetestsadmissions {
	return c.service.FindByID(id)
}

//FindByProjectID method
func (c *controller) FindByTestTypeID(id string) []entity.Basetestsadmissions {
	return c.service.FindByTestTypeID(id)
}
