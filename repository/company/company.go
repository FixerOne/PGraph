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

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Company{}, &entity.City{}, &entity.DataState{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Company) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Company) {}

func (r *repository) Delete(data entity.Company) {}

func (r *repository) FindAll() []entity.Company {
	var companies []entity.Company
	r.connection.Set("gorm:auto_preload", true).Find(&companies)
	return companies
}

func (r *repository) FindByID(id string) entity.Company {
	var company entity.Company
	r.connection.Set("gorm:auto_preload", true).First(&company, id)
	//r.connection.Set("gorm:auto_preload", true).Where("id = ?", "13").First(&company)
	return company
}
