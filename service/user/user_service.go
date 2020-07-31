package user

import (
	"pgraph/entity"
	repository "pgraph/repository/user"
)

//Service inerface
type Service interface {
	Save(entity.Company) entity.Company
	FindAll() []entity.User
	FindByID(id string) entity.Company
	FindByMail(entity.User) entity.User
	Login(entity.User) entity.User
	FindByCompanyID(id string) []entity.User
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
func (service *service) FindAll() []entity.User {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Company {
	return service.repository.FindByID(id)
}

//FindById Meethod
func (service *service) FindByCompanyID(id string) []entity.User {
	return service.repository.FindByCompanyID(id)
}

func (service *service) FindByMail(user entity.User) entity.User {
	return service.repository.FindByMail(user)
}

func (service *service) Login(user entity.User) entity.User {
	return service.repository.Login(user)
}
