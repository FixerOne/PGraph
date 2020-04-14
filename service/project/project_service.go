package project

import (
	"pgraph/entity"
	repository "pgraph/repository/project"
)

//Service inerface
type Service interface {
	Save(entity.Project) entity.Project
	FindAll() []entity.Project
	FindByID(id string) entity.Project
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
func (service *service) Save(dataEntity entity.Project) entity.Project {
	//service.companies = append(service.companies, company)
	service.repository.Save(dataEntity)
	return dataEntity
}

//FindAll Meethod
func (service *service) FindAll() []entity.Project {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Project {
	return service.repository.FindByID(id)
}
