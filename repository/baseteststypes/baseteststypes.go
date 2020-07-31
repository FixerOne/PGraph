package baseteststypes

import (
	"pgraph/database"
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
	return &repository{}
}

func (r *repository) Save(data entity.Baseteststypes) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Baseteststypes{})

	db.Save(&data)

	defer db.Close()
}

func (r *repository) Update(data entity.Baseteststypes) {}

func (r *repository) Delete(data entity.Baseteststypes) {}

func (r *repository) FindAll() []entity.Baseteststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Baseteststypes{})

	var entitites []entity.Baseteststypes

	db.Set("gorm:auto_preload", true).Find(&entitites)

	defer db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Baseteststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Baseteststypes{})

	var data entity.Baseteststypes
	db.Set("gorm:auto_preload", true).First(&data, id)

	db.Close()

	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Baseteststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Baseteststypes{})

	var data []entity.Baseteststypes
	db.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)

	defer db.Close()

	return data
}
