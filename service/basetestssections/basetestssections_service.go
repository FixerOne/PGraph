package basetestssections

import (
	"pgraph/entity"
	repository "pgraph/repository/basetestssections"
)

//Service inerface
type Service interface {
	Save(entity.Basetestssections) entity.Basetestssections
	FindAll() []entity.Basetestssections
	FindByID(id string) entity.Basetestssections
	FindByTestTypeID(id string) []entity.Basetestssections
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
func (service *service) Save(data entity.Basetestssections) entity.Basetestssections {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Basetestssections {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Basetestssections {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Basetestssections {
	return service.repository.FindByTestTypeID(id)
}
