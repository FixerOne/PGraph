package testtype

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Teststypes)
	Update(data entity.Teststypes)
	Delete(data entity.Teststypes)
	FindAll() []entity.Teststypes
	FindByID(id string) entity.Teststypes
	FindByProjectID(id string) []entity.Teststypes
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Teststypes) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Teststypes{})

	db.Save(&data)

	db.Close()

}

func (r *repository) Update(data entity.Teststypes) {}

func (r *repository) Delete(data entity.Teststypes) {}

func (r *repository) FindAll() []entity.Teststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Teststypes{})

	var entitites []entity.Teststypes

	db.Set("gorm:auto_preload", true).Find(&entitites)

	db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Teststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Teststypes{})

	var data entity.Teststypes
	db.Set("gorm:auto_preload", true).First(&data, id)

	db.Close()

	return data
}

func (r *repository) FindByProjectID(id string) []entity.Teststypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Teststypes{})

	var data []entity.Teststypes
	db.Set("gorm:auto_preload", true).Where("projects_id = ?", id).Find(&data)

	db.Close()

	return data
}
