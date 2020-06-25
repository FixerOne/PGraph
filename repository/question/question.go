package question

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Questions)
	Update(data entity.Questions)
	Delete(data entity.Questions)
	FindAll() []entity.Questions
	FindByID(id string) entity.Questions
	FindByTestTypeID(id string) []entity.Questions
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Questions{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Questions) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Questions) {}

func (r *repository) Delete(data entity.Questions) {}

func (r *repository) FindAll() []entity.Questions {
	var entitites []entity.Questions
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Questions {
	var data entity.Questions
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Questions {
	var data []entity.Questions
	//r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_projects_by_company_id(?);", id).Scan(&data)
	r.connection.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)
	return data
}
