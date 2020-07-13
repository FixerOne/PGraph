package basetestssections

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Basetestssections)
	Update(data entity.Basetestssections)
	Delete(data entity.Basetestssections)
	FindAll() []entity.Basetestssections
	FindByID(id string) entity.Basetestssections
	FindByTestTypeID(id string) []entity.Basetestssections
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestssections{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Basetestssections) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Basetestssections) {}

func (r *repository) Delete(data entity.Basetestssections) {}

func (r *repository) FindAll() []entity.Basetestssections {
	var entitites []entity.Basetestssections
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Basetestssections {
	var data entity.Basetestssections
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Basetestssections {
	var data []entity.Basetestssections
	r.connection.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)
	return data
}
