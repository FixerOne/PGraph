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
	FindByMail(data entity.User) entity.User
	Login(data entity.User) entity.User
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {
	return &repository{}
}

func (r *repository) Save(data entity.Company) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	db.Save(&data)

	db.Close()

}

func (r *repository) Update(data entity.Company) {}

func (r *repository) Delete(data entity.Company) {}

func (r *repository) FindAll() []entity.User {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	var entitites []entity.User
	db.Set("gorm:auto_preload", true).Order("first_name asc").Find(&entitites)

	db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Company {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	var company entity.Company
	db.Set("gorm:auto_preload", true).First(&company, id)

	db.Close()

	return company
}

func (r *repository) FindByCompanyID(id string) []entity.User {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	var data []entity.User
	db.Set("gorm:auto_preload", true).Where("company_id = ?", id).Find(&data)

	db.Close()

	return data
}

func (r *repository) FindByMail(data entity.User) entity.User {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	db.Set("gorm:auto_preload", true).Where("mail = ?", data.Mail).Find(&data)

	db.Close()

	return data
}

func (r *repository) Login(data entity.User) entity.User {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{}, &entity.Company{}, &entity.City{}, &entity.DataState{})

	db.Set("gorm:auto_preload", true).Where("mail = ? AND password = ?", data.Mail, data.Password).Find(&data)

	db.Close()

	return data
}
