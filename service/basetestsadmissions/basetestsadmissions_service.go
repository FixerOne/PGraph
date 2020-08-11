package basetestsadmissions

import (
	"pgraph/entity"
	repository "pgraph/repository/basetestsadmissions"
)

//Service inerface
type Service interface {
	Save(entity.Basetestsadmissions) entity.Basetestsadmissions
	FindAll() []entity.Basetestsadmissions
	FindByID(id string) entity.Basetestsadmissions
	FindByTestTypeID(id string) []entity.Basetestsadmissions
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
func (service *service) Save(data entity.Basetestsadmissions) entity.Basetestsadmissions {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Basetestsadmissions {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Basetestsadmissions {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Basetestsadmissions {
	return service.repository.FindByTestTypeID(id)
}
