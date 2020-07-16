package location

import (
	database "pgraph/database"
	"pgraph/entity"

	"github.com/jinzhu/gorm"
)

//Repository company
type Repository interface {
	Save(data entity.Country)
	UpdateCountry(data entity.Country)
	//Delete(data entity.Country)
	FindAllCountries() []entity.Country
	FindStatesByCountry(id string) []entity.State
	FindAllStates() []entity.State
	FindCitiesByCountry(id string) []entity.City
	FindCitiesByState(id string) []entity.City
	FindAllCities() []entity.City
}

type repository struct {
	connection *gorm.DB
}

//New Constructor
func New() Repository {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	return &repository{
		connection: db,
	}
}

func (r *repository) Save(data entity.Country) {
	r.connection.Save(&data)
}

func (r *repository) UpdateCountry(data entity.Country) {
	r.connection.Save(&data)
}

func (r *repository) FindAllCountries() []entity.Country {
	var data []entity.Country
	r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&data)
	return data
}

func (r *repository) FindAllStates() []entity.State {
	var data []entity.State
	r.connection.Set("gorm:auto_preload", true).Order("country_id asc").Find(&data)
	return data
}

func (r *repository) FindStatesByCountry(id string) []entity.State {
	var data []entity.State
	r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_states_by_country_id(?);", id).Scan(&data)
	return data
}

func (r *repository) FindCitiesByCountry(id string) []entity.City {
	var data []entity.City
	r.connection.Set("gorm:auto_preload", true).Joins("JOIN states on states.id=cities.state_id").Where("country_id = ?", id).Find(&data)
	return data
}

func (r *repository) FindCitiesByState(id string) []entity.City {
	var data []entity.City
	r.connection.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_cities_by_state_id(?);", id).Scan(&data)
	return data
}

func (r *repository) FindAllCities() []entity.City {
	var data []entity.City
	r.connection.Set("gorm:auto_preload", true).Order("name asc").Find(&data)
	return data
}
