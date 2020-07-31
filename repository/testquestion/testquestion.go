package testquestion

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Testquestion)
	Update(data entity.Testquestion)
	Delete(data entity.Testquestion)
	FindAll() []entity.Testquestion
	FindByID(id string) entity.Testquestion
	FindByTestTypeID(id string) []entity.Testquestion
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Testquestion) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Testquestion) {}

func (r *repository) Delete(data entity.Testquestion) {}

func (r *repository) FindAll() []entity.Testquestion {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Testquestion{}, &entity.Teststypes{}, &entity.Basetestsquestions{})

	var entitites []entity.Testquestion

	db.Set("gorm:auto_preload", true).Find(&entitites)

	db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Testquestion {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Testquestion{}, &entity.Teststypes{}, &entity.Basetestsquestions{})

	var data entity.Testquestion
	db.Set("gorm:auto_preload", true).First(&data, id)

	db.Close()

	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Testquestion {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Testquestion{}, &entity.Teststypes{}, &entity.Basetestsquestions{})

	var data []entity.Testquestion

	db.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)

	db.Close()

	return data
}
