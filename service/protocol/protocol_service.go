package protocol

import (
	"pgraph/entity"
	repository "pgraph/repository/project"
)

//Service inerface
type Service interface {
	Save(entity.Project) entity.Protocol
	FindAll() []entity.Protocol
	FindByID(id string) entity.Protocol
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
func (service *service) Save(dataEntity entity.Protocol) entity.Protocol {
	//service.companies = append(service.companies, company)
	service.repository.Save(dataEntity)
	return dataEntity
}

//FindAll Meethod
func (service *service) FindAll() []entity.Protocol {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Protocol {
	return service.repository.FindByID(id)
}
