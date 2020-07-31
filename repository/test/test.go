package test

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Test)
	Update(data entity.Test)
	Delete(data entity.Test)
	FindAll() []entity.Test
	FindByID(id string) entity.Test
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Test{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Test) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Test) {
	r.connection.Save(&data)
}

func (r *repository) Delete(data entity.Test) {}

func (r *repository) FindAll() []entity.Test {
	var data []entity.Test
	r.connection.Set("gorm:auto_preload", true).Order("id asc").Find(&data)
	return data
}

func (r *repository) FindByID(id string) entity.Test {
	var dataEntity entity.Test
	r.connection.Set("gorm:auto_preload", true).First(&dataEntity, id)
	return dataEntity
}
