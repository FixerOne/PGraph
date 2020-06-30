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

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Basetestsquestions{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Basetestsquestions) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Basetestsquestions) {}

func (r *repository) Delete(data entity.Basetestsquestions) {}

func (r *repository) FindAll() []entity.Basetestsquestions {
	var entitites []entity.Basetestsquestions
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Basetestsquestions {
	var data entity.Basetestsquestions
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByTestTypeID(id string) []entity.Basetestsquestions {
	var data []entity.Basetestsquestions
	//r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_projects_by_company_id(?);", id).Scan(&data)
	r.connection.Set("gorm:auto_preload", true).Where("teststypes_id = ?", id).Find(&data)
	return data
}
