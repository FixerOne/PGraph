package baseteststypes

import (
	"pgraph/entity"
	repository "pgraph/repository/baseteststypes"
)

//Service inerface
type Service interface {
	Save(entity.Baseteststypes) entity.Baseteststypes
	FindAll() []entity.Baseteststypes
	FindByID(id string) entity.Baseteststypes
	FindByTestTypeID(id string) []entity.Baseteststypes
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
func (service *service) Save(data entity.Baseteststypes) entity.Baseteststypes {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Baseteststypes {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Baseteststypes {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Baseteststypes {
	return service.repository.FindByTestTypeID(id)
}
