package location

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Country)
	Update(data entity.Country)
	//Delete(data entity.Country)
	FindAll() []entity.Country
}

type repository struct {
	connection *gorm.DB
}

//New Constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Country) {
	r.connection.Save(&data)
}

func (r *repository) Update(data entity.Country) {}

//func (r *repository) Delete(data entity.Country) {}

func (r *repository) FindAll() []entity.Country {
	var data []entity.Country
	r.connection.Set("gorm:auto_preload", true).Find(&data)
	return data
}
