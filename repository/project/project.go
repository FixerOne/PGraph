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
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Company{}, &entity.City{}, &entity.DataState{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Project) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Project) {}

func (r *repository) Delete(data entity.Project) {}

func (r *repository) FindAll() []entity.Project {
	var entitites []entity.Project
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Project {
	var dataEntity entity.Project
	r.connection.Set("gorm:auto_preload", true).First(&dataEntity, id)
	return dataEntity
}
