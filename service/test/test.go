package test

import (
	"pgraph/entity"
	repository "pgraph/repository/test"
)

//Service inerface
type Service interface {
	Save(entity.Test) entity.Test
	Update(entity.Test) entity.Test
	FindAll() []entity.Test
	FindByID(id string) entity.Test
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
func (service *service) Save(dataEntity entity.Test) entity.Test {
	//service.companies = append(service.companies, company)
	service.repository.Save(dataEntity)
	return dataEntity
}

//Update method
func (service *service) Update(dataEntity entity.Test) entity.Test {
	//service.companies = append(service.companies, company)
	service.repository.Update(dataEntity)
	return dataEntity
}

//FindAll Meethod
func (service *service) FindAll() []entity.Test {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Test {
	return service.repository.FindByID(id)
}
