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
	return &repository{}
}

func (r *repository) Save(data entity.Basetestssections) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestssections{})

	db.Save(&data)

	defer db.Close()

}

func (r *repository) Update(data entity.Basetestssections) {}

func (r *repository) Delete(data entity.Basetestssections) {}

func (r *repository) FindAll() []entity.Basetestssections {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestssections{})

	var entitites []entity.Basetestssections

	db.Set("gorm:auto_preload", true).Find(&entitites)

	defer db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Basetestssections {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestssections{})

	var data entity.Basetestssections
	db.Set("gorm:auto_preload", true).First(&data, id)

	defer db.Close()

	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Basetestssections {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestssections{})

	var data []entity.Basetestssections
	db.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)

	defer db.Close()

	return data
}
