package user

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
	FindAll() []entity.User
	FindByID(id string) entity.Company
	FindByCompanyID(id string) []entity.User
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Company) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Company) {}

func (r *repository) Delete(data entity.Company) {}

func (r *repository) FindAll() []entity.User {
	var entitites []entity.User
	r.connection.Set("gorm:auto_preload", true).Order("first_name asc").Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Company {
	var company entity.Company
	r.connection.Set("gorm:auto_preload", true).First(&company, id)
	return company
}

func (r *repository) FindByCompanyID(id string) []entity.User {
	var data []entity.User
	r.connection.Set("gorm:auto_preload", true).Where("company_id = ?", id).Find(&data)
	return data
}
