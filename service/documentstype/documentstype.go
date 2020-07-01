package documentstype

import (
	"pgraph/entity"
	repository "pgraph/repository/documentstype"
)

//Service inerface
type Service interface {
	Save(entity.Documentstypes) entity.Documentstypes
	FindAll() []entity.Documentstypes
	FindByID(id string) entity.Documentstypes
	FindByProjectID(id string) []entity.Documentstypes
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
func (service *service) Save(data entity.Documentstypes) entity.Documentstypes {
	service.repository.Save(data)
	return data
}

//FindAll Meethod
func (service *service) FindAll() []entity.Documentstypes {
	return service.repository.FindAll()
}

//FindById Meethod
func (service *service) FindByID(id string) entity.Documentstypes {
	return service.repository.FindByID(id)
}

//FindByProjectID Meethod
func (service *service) FindByProjectID(id string) []entity.Documentstypes {
	return service.repository.FindByProjectID(id)
}
