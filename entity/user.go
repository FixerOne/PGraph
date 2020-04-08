package entity

//User user info
type User struct {
	ID                int16     `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	FirstName         string    `json:"first_name,omitempty" binding:"required" gorm:"type:varchar(50)"`
	LastName          string    `json:"last_name,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Birthday          string    `json:"Birthday,omitempty"`
	Document          string    `json:"Document,omitempty"`
	DocumentType      string    `json:"DocumentType,omitempty"`
	Address           string    `json:"Address,omitempty"`
	Gender            string    `json:"Gender,omitempty"`
	CompanyID         string    `json:"company_id,omitempty"`
	Phone             string    `json:"Phone,omitempty"`
	Email             string    `json:"Email,omitempty"`
	UserType          string    `json:"UserType,omitempty"`
	Created           string    `json:"Created,omitempty"`
	Updated           string    `json:"Updated,omitempty"`
	State             string    `json:"State,omitempty"`
	Password          string    `json:"Password,omitempty"`
	DataState         DataState `json:"DataState" gorm:"foreignkey:data_state_id"`
	DataStateID       int       `json:"DataStateID,omitempty"`
	CityNationality   City      `json:"CityNationality" gorm:"foreignkey:nationality_city_id"`
	CityNationalityID int       `json:"nationality_city_id,omitempty"`
	CityResidence     City      `json:"CityResidence" gorm:"foreignkey:residence_city_id"`
	CityResidenceID   int       `json:"CityResidenceID,omitempty"`

	//nationality_city_id_fk
}
