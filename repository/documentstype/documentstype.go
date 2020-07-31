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
	return &repository{}
}

func (r *repository) Save(data entity.Documentstypes) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{})

	db.Save(&data)

	db.Close()

}

func (r *repository) Update(data entity.Documentstypes) {}

func (r *repository) Delete(data entity.Documentstypes) {}

func (r *repository) FindAll() []entity.Documentstypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{})

	var entitites []entity.Documentstypes
	db.Set("gorm:auto_preload", true).Find(&entitites)

	defer db.Close()

	return entitites
}

func (r *repository) FindByID(id string) entity.Documentstypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{})

	var data entity.Documentstypes
	db.Set("gorm:auto_preload", true).First(&data, id)

	db.Close()

	return data
}

func (r *repository) FindByProjectID(id string) []entity.Documentstypes {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Documentstypes{})

	var data []entity.Documentstypes
	db.Set("gorm:auto_preload", true).Where("projects_id = ?", id).Find(&data)

	defer db.Close()

	return data
}
