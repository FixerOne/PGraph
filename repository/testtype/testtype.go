package testtype

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Teststypes)
	Update(data entity.Teststypes)
	Delete(data entity.Teststypes)
	FindAll() []entity.Teststypes
	FindByID(id string) entity.Teststypes
	FindByProjectID(id string) []entity.Teststypes
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Teststypes{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Teststypes) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Teststypes) {}

func (r *repository) Delete(data entity.Teststypes) {}

func (r *repository) FindAll() []entity.Teststypes {
	var entitites []entity.Teststypes
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Teststypes {
	var data entity.Teststypes
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByProjectID(id string) []entity.Teststypes {
	var data []entity.Teststypes
	//r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_projects_by_company_id(?);", id).Scan(&data)
	r.connection.Set("gorm:auto_preload", true).Where("projects_id = ?", id).Find(&data)
	return data
}
