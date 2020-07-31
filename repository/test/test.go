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
	return &repository{}
}

func (r *repository) Save(data entity.Test) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Test{})

	db.Save(&data)

	db.Close()
}

func (r *repository) Update(data entity.Test) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Test{})

	db.Save(&data)

	db.Close()
}

func (r *repository) Delete(data entity.Test) {}

func (r *repository) FindAll() []entity.Test {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Test{})

	var data []entity.Test
	db.Set("gorm:auto_preload", true).Order("id asc").Find(&data)

	db.Close()

	return data
}

func (r *repository) FindByID(id string) entity.Test {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Test{})

	var dataEntity entity.Test
	db.Set("gorm:auto_preload", true).First(&dataEntity, id)

	db.Close()

	return dataEntity
}
