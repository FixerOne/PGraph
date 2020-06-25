package testquestion

import (
	"pgraph/entity"
	repository "pgraph/repository/testquestion"
)

//Service inerface
type Service interface {
	Save(entity.Testquestion) entity.Testquestion
	FindAll() []entity.Testquestion
	FindByID(id string) entity.Testquestion
	FindByTestTypeID(id string) []entity.Testquestion
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
func (service *service) Save(data entity.Testquestion) entity.Testquestion {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Testquestion {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Testquestion {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Testquestion {
	return service.repository.FindByTestTypeID(id)
}
