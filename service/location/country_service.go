package location

import (
	"pgraph/entity"
	repository "pgraph/repository/location"
)

//Service inerface
type Service interface {
	Save(entity.Country) entity.Country
	FindAll() []entity.Country
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
func (service *service) Save(data entity.Country) entity.Country {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Country {
	return service.repository.FindAll()
}
