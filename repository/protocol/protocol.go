package protocol

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Protocol)
	Update(data entity.Protocol)
	Delete(data entity.Protocol)
	FindAll() []entity.Protocol
	FindByID(id string) entity.Protocol
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Protocol{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Protocol) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Protocol) {}

func (r *repository) Delete(data entity.Protocol) {}

func (r *repository) FindAll() []entity.Protocol {
	var entitites []entity.Protocol
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Protocol {
	var dataEntity entity.Protocol
	r.connection.Set("gorm:auto_preload", true).First(&dataEntity, id)
	return dataEntity
}
