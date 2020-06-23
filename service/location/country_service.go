package location

import (
	"pgraph/entity"
	repository "pgraph/repository/location"
)

//Service inerface
type Service interface {
	Save(entity.Country) entity.Country
	FindAllCountries() []entity.Country
	FindStatesByCountry(id string) []entity.State
	FindCitiesByState(id string) []entity.City
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
func (service *service) FindAllCountries() []entity.Country {
	return service.repository.FindAllCountries()
}

//FindAll Meethod
func (service *service) FindStatesByCountry(id string) []entity.State {
	return service.repository.FindStatesByCountry(id)
}

func (service *service) FindCitiesByState(id string) []entity.City {
	return service.repository.FindCitiesByState(id)
}
