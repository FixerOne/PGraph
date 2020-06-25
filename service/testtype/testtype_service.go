package testtype

import (
	"pgraph/entity"
	repository "pgraph/repository/testtype"
)

//Service inerface
type Service interface {
	Save(entity.Teststypes) entity.Teststypes
	FindAll() []entity.Teststypes
	FindByID(id string) entity.Teststypes
	FindByProjectID(id string) []entity.Teststypes
}

type service struct {
	repository repository.Repository
}

//New returns service instance
func New(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

//Save method
func (service *service) Save(data entity.Teststypes) entity.Teststypes {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Teststypes {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Teststypes {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByProjectID(id string) []entity.Teststypes {
	return service.repository.FindByProjectID(id)
}
