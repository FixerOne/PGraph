package documentstype

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Documentstypes)
	Update(data entity.Documentstypes)
	Delete(data entity.Documentstypes)
	FindAll() []entity.Documentstypes
	FindByID(id string) entity.Documentstypes
	FindByProjectID(id string) []entity.Documentstypes
}

type repository struct {
	connection *gorm.DB
}

//New constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Documentstypes) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Documentstypes) {}

func (r *repository) Delete(data entity.Documentstypes) {}

func (r *repository) FindAll() []entity.Documentstypes {
	var entitites []entity.Documentstypes
	//r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&entitites)
	r.connection.Set("gorm:auto_preload", true).Find(&entitites)
	return entitites
}

func (r *repository) FindByID(id string) entity.Documentstypes {
	var data entity.Documentstypes
	r.connection.Set("gorm:auto_preload", true).First(&data, id)
	return data
}

func (r *repository) FindByProjectID(id string) []entity.Documentstypes {
	var data []entity.Documentstypes
	//r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_projects_by_company_id(?);", id).Scan(&data)
	r.connection.Set("gorm:auto_preload", true).Where("projects_id = ?", id).Find(&data)
	return data
}
