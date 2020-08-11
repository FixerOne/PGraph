package entity

//User user info
type User struct {
	ID                int16          `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	FirstName         string         `json:"FirstName,omitempty" binding:"required" gorm:"type:varchar(50)"`
	LastName          string         `json:"LastName,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Company           Company        `json:"Company" gorm:"foreignkey:company_id"`
	CompanyID         string         `json:"company_id,omitempty"`
	Birthday          string         `json:"Birthday,omitempty" binding:"required" gorm:"type:date(100)"`
	Document          string         `json:"Document,omitempty"`
	DocumentType      Documentstypes `json:"DocumentType" gorm:"foreignkey:documentstypes_id"`
	DocumentstypesID  int            `json:"documentstypes_id,omitempty"`
	Address           string         `json:"Address,omitempty" binding:"required" gorm:"type:varchar(200)"`
	Phone             string         `json:"Phone,omitempty" binding:"required" gorm:"type:varchar(20)"`
	NationalityCity   City           `json:"CityNationality" gorm:"foreignkey:nationality_city_id"`
	NationalityCityID int            `json:"nationality_city_id,omitempty"`
	Gender            string         `json:"Gender,omitempty" binding:"required" gorm:"type:varchar(1)"`
	ResidenceCity     City           `json:"CityResidence" gorm:"foreignkey:residence_city_id"`
	ResidenceCityID   int            `json:"CityResidenceID,omitempty"`
	Created           string         `json:"Created" gorm:"default:CURRENT_TIMESTAMP"`
	Updated           string         `json:"Updated" gorm:"default:CURRENT_TIMESTAMP"`
	DataState         DataState      `json:"DataState" gorm:"foreignkey:data_state_id"`
	DataStateID       int            `json:"DataStateID,omitempty"`
	PhoneContact      string         `json:"PhoneContact,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Mail              string         `json:"Mail,omitempty" binding:"required" gorm:"type:varchar(256)"`
	Password          string         `json:"Password,omitempty" binding:"required" gorm:"type:varchar(256)"`
	Userstypes        Userstypes     `json:"Userstype" gorm:"foreignkey:userstypes_id"`
	UserstypesID      string         `json:"userstypes_id,omitempty" binding:"required" gorm:"type:int4"`
	Token             string         `json:"Token,omitempty"`
}
