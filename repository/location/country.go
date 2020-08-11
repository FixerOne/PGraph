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
	FindAllCountriesByActive() []entity.Country
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
	return &repository{}
}

func (r *repository) Save(data entity.Country) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	db.Save(&data)

	db.Close()

}

func (r *repository) UpdateCountry(data entity.Country) {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	db.Save(&data)

	db.Close()

}

func (r *repository) FindAllCountries() []entity.Country {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.Country
	db.Set("gorm:auto_preload", true).Order("name asc").Find(&data)

	db.Close()

	return data
}

func (r *repository) FindAllCountriesByActive() []entity.Country {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.Country
	db.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_countries_by_active(?);", true).Scan(&data)

	db.Close()

	return data
}

func (r *repository) FindAllStates() []entity.State {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.State
	db.Set("gorm:auto_preload", true).Order("country_id asc").Find(&data)

	db.Close()

	return data
}

func (r *repository) FindStatesByCountry(id string) []entity.State {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.State
	db.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_states_by_country_id(?);", id).Scan(&data)

	db.Close()

	return data
}

func (r *repository) FindCitiesByCountry(id string) []entity.City {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.City
	db.Set("gorm:auto_preload", true).Joins("JOIN states on states.id=cities.state_id").Where("country_id = ?", id).Find(&data)

	db.Close()

	return data
}

func (r *repository) FindCitiesByState(id string) []entity.City {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.City
	db.Set("gorm:auto_preload", true).Raw("SELECT * from public.get_cities_by_state_id(?);", id).Scan(&data)

	db.Close()

	return data
}

func (r *repository) FindAllCities() []entity.City {

	database.Init()
	db := database.GetDB()
	db.AutoMigrate(&entity.Country{}, &entity.State{}, &entity.City{})

	var data []entity.City
	db.Set("gorm:auto_preload", true).Order("name asc").Find(&data)

	db.Close()

	return data
}
