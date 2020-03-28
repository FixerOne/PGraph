package entity

//Company company info
type Company struct {
	ID           int       `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name         string    `json:"Name,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Phone        string    `json:"Phone,omitempty" binding:"required" gorm:"type:varchar(20)"`
	Address      string    `json:"Address,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Contact      string    `json:"Contact,omitempty" binding:"required" gorm:"type:varchar(50)"`
	PhoneContact string    `json:"PhoneContact,omitempty" binding:"required" gorm:"type:varchar(50)"`
	MailContact  string    `json:"MailContact,omitempty" binding:"required" gorm:"type:varchar(256)"`
	NIT          string    `json:"NIT,omitempty" binding:"required" gorm:"type:varchar(20)"`
	Created      string    `json:"Created" gorm:"default:CURRENT_TIMESTAMP"`
	Updated      string    `json:"Updated" gorm:"default:CURRENT_TIMESTAMP"`
	DataState    DataState `json:"DataState" gorm:"foreignkey:data_state_id"`
	DataStateID  int       `json:"DataStateID,omitempty"`
	City         City      `json:"City" gorm:"foreignkey:city_id"`
	CityID       int       `json:"CityID,omitempty"`
}

//City data
type City struct {
	ID      int    `gorm:"primary_key" json:"ID,omitempty"`
	Name    string `json:"name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	State   State  `json:"State" gorm:"foreignkey:state_id"`
	StateID string `json:"StateID,omitempty"`
}

//State data
type State struct {
	ID        int     `gorm:"primary_key" json:"ID,omitempty"`
	Name      string  `json:"name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Country   Country `json:"Country" gorm:"foreignkey:country_id"`
	CountryID string  `json:"CountryID,omitempty"`
}

//Country data
type Country struct {
	ID        int    `gorm:"primary_key" json:"ID,omitempty"`
	ShortName string `json:"ShortName,omitempty" binding:"required" gorm:"type:varchar(2)"`
	Name      string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(150)"`
	PhoneCode int    `json:"PhoneCode,omitempty" binding:"required"`
}

//DataState data
type DataState struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(100)"`
}
