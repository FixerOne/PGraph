package baseteststypes

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Baseteststypes)
	Update(data entity.Baseteststypes)
	Delete(data entity.Baseteststypes)
	FindAll() []entity.Baseteststypes
	FindByID(id string) entity.Baseteststypes
	FindByTestTypeID(id string) []entity.Baseteststypes
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Baseteststypes{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Baseteststypes) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Baseteststypes) {}

func (r *repository) Delete(data entity.Baseteststypes) {}

func (r *repository) FindAll() []entity.Baseteststypes {
	var entitites []entity.Baseteststypes
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Baseteststypes {
	var data entity.Baseteststypes
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Baseteststypes {
	var data []entity.Baseteststypes
	r.connection.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)
	return data
}
