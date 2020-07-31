package basequestion

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Basetestsquestions)
	Update(data entity.Basetestsquestions)
	Delete(data entity.Basetestsquestions)
	FindAll() []entity.Basetestsquestions
	FindByID(id string) entity.Basetestsquestions
	FindByTestTypeID(id string) []entity.Basetestsquestions
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Basetestsquestions) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Basetestsquestions) {
	r.connection.Save(&data)
}

func (r *repository) Delete(data entity.Basetestsquestions) {}

func (r *repository) FindAll() []entity.Basetestsquestions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsquestions{})

	var entitites []entity.Basetestsquestions

	db.Set("gorm:auto_preload", true).Find(&entitites)
	defer db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Basetestsquestions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsquestions{})

	var data entity.Basetestsquestions
	db.Set("gorm:auto_preload", true).First(&data, id)

	defer db.Close()

	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Basetestsquestions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsquestions{})

	var data []entity.Basetestsquestions
	db.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)

	defer db.Close()

	return data
}
