package company

import (
	"pgraph/entity"
	repository "pgraph/repository/company"
)

//CompanyService inerface
type Service interface {
	Save(entity.Company) entity.Company
	FindAll() []entity.Company
	FindByID(id string) entity.Company
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
func (service *service) Save(company entity.Company) entity.Company {
	//service.companies = append(service.companies, company)
	service.repository.Save(company)
	return company
}

//FindAll Meethod
func (service *service) FindAll() []entity.Company {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Company {
	return service.repository.FindByID(id)
}
