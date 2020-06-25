package question

import (
	"pgraph/entity"
	repository "pgraph/repository/question"
)

//Service inerface
type Service interface {
	Save(entity.Questions) entity.Questions
	FindAll() []entity.Questions
	FindByID(id string) entity.Questions
	FindByTestTypeID(id string) []entity.Questions
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
func (service *service) Save(data entity.Questions) entity.Questions {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Questions {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Questions {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Questions {
	return service.repository.FindByTestTypeID(id)
}
