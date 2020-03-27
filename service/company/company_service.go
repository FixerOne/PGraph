package company

import (
	"pgraph/entity"
	repository "pgraph/repository/company"
)

//CompanyService inerface
type CompanyService interface {
	Save(entity.Company) entity.Company
	FindAll() []entity.Company
}

type companyService struct {
	repository repository.Repository
}

//New returns service instance
func New(repository repository.Repository) CompanyService {
	return &companyService{
		repository: repository,
	}
}

//Save method
func (service *companyService) Save(company entity.Company) entity.Company {
	//service.companies = append(service.companies, company)
	service.repository.Save(company)
	return company
}

//FindAll Meethod
func (service *companyService) FindAll() []entity.Company {
	return service.repository.FindAll()
}
