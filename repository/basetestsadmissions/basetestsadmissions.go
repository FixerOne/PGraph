package basetestsadmissions

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Basetestsadmissions)
	Update(data entity.Basetestsadmissions)
	Delete(data entity.Basetestsadmissions)
	FindAll() []entity.Basetestsadmissions
	FindByID(id string) entity.Basetestsadmissions
	FindByTestTypeID(id string) []entity.Basetestsadmissions
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Basetestsadmissions) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsadmissions{})

	db.Save(&data)

	defer db.Close()

}

func (r *repository) Update(data entity.Basetestsadmissions) {}

func (r *repository) Delete(data entity.Basetestsadmissions) {}

func (r *repository) FindAll() []entity.Basetestsadmissions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsadmissions{})

	var entitites []entity.Basetestsadmissions

	db.Set("gorm:auto_preload", true).Find(&entitites)

	defer db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Basetestsadmissions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsadmissions{})

	var data entity.Basetestsadmissions
	db.Set("gorm:auto_preload", true).First(&data, id)

	defer db.Close()

	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Basetestsadmissions {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsadmissions{})

	var data []entity.Basetestsadmissions
	db.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)

	defer db.Close()

	return data
}
