package company

import (
	"pgraph/entity"
	service "pgraph/service/company"

	"github.com/gin-gonic/gin"
)

//CompanyController inerface
type CompanyController interface {
	FindAll() []entity.Company
	Save(ctx *gin.Context) entity.Company
}

type controller struct {
	service service.CompanyService
}

//New constructor
func New(service service.CompanyService) CompanyController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Company {
	var company entity.Company
	ctx.BindJSON(&company)
	return c.service.Save(company)
}

//FindAll method
func (c *controller) FindAll() []entity.Company {
	return c.service.FindAll()
}
