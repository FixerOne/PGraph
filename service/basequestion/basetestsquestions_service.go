package basequestion

import (
	"pgraph/entity"
	repository "pgraph/repository/basequestion"
)

//Service inerface
type Service interface {
	Save(entity.Basetestsquestions) entity.Basetestsquestions
	FindAll() []entity.Basetestsquestions
	FindByID(id string) entity.Basetestsquestions
	FindByTestTypeID(id string) []entity.Basetestsquestions
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
func (service *service) Save(data entity.Basetestsquestions) entity.Basetestsquestions {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Basetestsquestions {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Basetestsquestions {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByTestTypeID(id string) []entity.Basetestsquestions {
	return service.repository.FindByTestTypeID(id)
}
