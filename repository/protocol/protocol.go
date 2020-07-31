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
	return &repository{}
}

func (r *repository) Save(data entity.Protocol) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Protocol{})

	db.Save(&data)

	db.Close()

}

func (r *repository) Update(data entity.Protocol) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Protocol{})

	db.Save(&data)

	db.Close()
}

func (r *repository) Delete(data entity.Protocol) {}

func (r *repository) FindAll() []entity.Protocol {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Protocol{})

	var data []entity.Protocol
	db.Set("gorm:auto_preload", true).Order("name asc").Find(&data)

	db.Close()

	return data
}

func (r *repository) FindByID(id string) entity.Protocol {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Protocol{})

	var dataEntity entity.Protocol
	db.Set("gorm:auto_preload", true).First(&dataEntity, id)

	db.Close()

	return dataEntity
}
