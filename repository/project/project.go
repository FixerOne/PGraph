package project

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Project)
	Update(data entity.Project)
	Delete(data entity.Project)
	FindAll() []entity.Project
	FindByID(id string) entity.Project
	FindByCompanyID(id string) []entity.Project
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Project) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(entity.Project{})

	db.Save(&data)

	db.Close()

}

func (r *repository) Update(data entity.Project) {}

func (r *repository) Delete(data entity.Project) {}

func (r *repository) FindAll() []entity.Project {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(entity.Project{})

	var entitites []entity.Project
	db.Set("gorm:auto_preload", true).Find(&entitites)

	db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Project {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(entity.Project{})

	var dataEntity entity.Project
	db.Set("gorm:auto_preload", true).First(&dataEntity, id)

	db.Close()

	return dataEntity
}

func (r *repository) FindByCompanyID(id string) []entity.Project {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(entity.Project{})

	var data []entity.Project
	db.Set("gorm:auto_preload", true).Where("company_id = ?", id).Find(&data)

	db.Close()

	return data
}
