package company

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Company)
	Update(data entity.Company)
	Delete(data entity.Company)
	FindAll() []entity.Company
	FindByID(id string) entity.Company
}

type repository struct {
	connection *gorm.DB
}

func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Company) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Company{}, &entity.City{}, &entity.DataState{})

	db.Save(&data)

	defer db.Close()

}

func (r *repository) Update(data entity.Company) {}

func (r *repository) Delete(data entity.Company) {}

func (r *repository) FindAll() []entity.Company {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Company{}, &entity.City{}, &entity.DataState{})

	var companies []entity.Company
	db.Set("gorm:auto_preload", true).Order("name asc").Find(&companies)

	defer db.Close()

	return companies
}

func (r *repository) FindByID(id string) entity.Company {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Company{}, &entity.City{}, &entity.DataState{})

	var company entity.Company
	db.Set("gorm:auto_preload", true).First(&company, id)

	defer db.Close()

	return company
}
