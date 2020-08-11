package user

import (
	"pgraph/entity"
	service "pgraph/service/user"

	"github.com/gin-gonic/gin"
)

//CompanyController inerface
type CompanyController interface {
	FindOne() entity.User
	FindByID(id string) entity.Company
	FindAll() []entity.User
	FindAllByType(usertype string) []entity.User
	FindByCompanyID(id string) []entity.User
	FindByMail(ctx *gin.Context) entity.User
	Login(ctx *gin.Context) entity.User
	Save(ctx *gin.Context) entity.Company
}

type controller struct {
	service service.Service
}

//New constructor
func New(service service.Service) CompanyController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.Company {
	var company entity.Company
	ctx.ShouldBindJSON(&company)
	return c.service.Save(company)
}

//FindAll method
func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

//FindAll method
func (c *controller) FindAllByType(usertype string) []entity.User {
	return c.service.FindAllByType(usertype)
}

//FindAll method
func (c *controller) FindOne() entity.User {
	return c.service.FindAll()[1]
}

//FindAll method
func (c *controller) FindByID(id string) entity.Company {
	return c.service.FindByID(id)
}

//FindAll method
func (c *controller) FindByCompanyID(id string) []entity.User {
	return c.service.FindByCompanyID(id)
}

func (c *controller) FindByMail(ctx *gin.Context) entity.User {
	var data entity.User
	ctx.ShouldBindJSON(&data)
	return c.service.FindByMail(data)
}

func (c *controller) Login(ctx *gin.Context) entity.User {
	var data entity.User
	ctx.ShouldBindJSON(&data)
	return c.service.Login(data)
}
