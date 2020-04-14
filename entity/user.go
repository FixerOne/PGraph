package entity

//User user info
type User struct {
	ID                int16     `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	FirstName         string    `json:"FirstName,omitempty" binding:"required" gorm:"type:varchar(50)"`
	LastName          string    `json:"LastName,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Company           Company   `json:"Company" gorm:"foreignkey:company_id"`
	CompanyID         string    `json:"company_id,omitempty"`
	Birthday          string    `json:"Birthday,omitempty" binding:"required" gorm:"type:date(100)"`
	Document          string    `json:"Document,omitempty"`
	DocumentType      string    `json:"DocumentType,omitempty"`
	Address           string    `json:"Address,omitempty" binding:"required" gorm:"type:varchar(200)"`
	Phone             string    `json:"Phone,omitempty" binding:"required" gorm:"type:varchar(20)"`
	NationalityCity   City      `json:"CityNationality" gorm:"foreignkey:nationality_city_id"`
	NationalityCityID int       `json:"nationality_city_id,omitempty"`
	Gender            string    `json:"Gender,omitempty" binding:"required" gorm:"type:varchar(1)"`
	ResidenceCity     City      `json:"CityResidence" gorm:"foreignkey:residence_city_id"`
	ResidenceCityID   int       `json:"CityResidenceID,omitempty"`
	Created           string    `json:"Created" gorm:"default:CURRENT_TIMESTAMP"`
	Updated           string    `json:"Updated" gorm:"default:CURRENT_TIMESTAMP"`
	DataState         DataState `json:"DataState" gorm:"foreignkey:data_state_id"`
	DataStateID       int       `json:"DataStateID,omitempty"`
	PhoneContact      string    `json:"PhoneContact,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Mail              string    `json:"MailContact,omitempty" binding:"required" gorm:"type:varchar(256)"`
	Password          string    `json:"Password,omitempty" binding:"required" gorm:"type:varchar(256)"`
	UserType          string    `json:"UserTypeId,omitempty" binding:"required" gorm:"type:int4"`
	//nationality_city_id_fk
}

/*
id OK
first_name OK
last_name OK
company_id OK
birthday OK
document OK
document_type OK
address OK
phone OK
nationality_city_id OK
gender OK
residence_city_id OK
created OK
updated OK
data_state_id OK
phone_contact
mail
password
user_type
*/
