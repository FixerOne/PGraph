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

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Testquestion{}, &entity.Teststypes{}, &entity.Basetestsquestions{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Testquestion) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Testquestion) {}

func (r *repository) Delete(data entity.Testquestion) {}

func (r *repository) FindAll() []entity.Testquestion {
	var entitites []entity.Testquestion
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Testquestion {
	var data entity.Testquestion
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Testquestion {
	var data []entity.Testquestion
	//r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_projects_by_company_id(?);", id).Scan(&data)
	r.connection.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)
	return data
}
