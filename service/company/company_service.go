package company

import "pgraph/entity"

//CompanyService inerface
type CompanyService interface {
	Save(entity.Company) entity.Company
	FindAll() []entity.Company
}

type companyService struct {
	companies []entity.Company
}

//New returns service instance
func New() CompanyService {
	return &companyService{}
}

//Save method
func (service *companyService) Save(company entity.Company) entity.Company {
	service.companies = append(service.companies, company)
	return company
}

//FindAll Meethod
func (service *companyService) FindAll() []entity.Company {
	return service.companies
}
