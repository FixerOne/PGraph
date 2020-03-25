package controller

import (
	"pgraph/entity"
	"pgraph/service"

	"github.com/gin-gonic/gin"
)

//UserController inerface
type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

//New constructor
func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

//Save method
func (c *controller) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.Save(user)
}

//FindAll method
func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}
