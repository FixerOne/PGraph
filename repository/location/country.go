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
	FindAllCountries() []entity.Country
	FindStatesByCountry(id string) []entity.State
	FindCitiesByState(id string) []entity.City
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

func (r *repository) FindAllCountries() []entity.Country {
	var data []entity.Country
	r.connection.Set("gorm:auto_preload", true).Find(&data)
	return data
}

func (r *repository) FindStatesByCountry(id string) []entity.State {
	var data []entity.State
	r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_states_by_country_id(?);", id).Scan(&data)
	return data
}

func (r *repository) FindCitiesByState(id string) []entity.City {
	var data []entity.City
	r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_cities_by_state_id(?);", id).Scan(&data)
	return data
}
