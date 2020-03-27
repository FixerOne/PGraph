package entity

//Company company info
type Company struct {
	ID           int       `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name         string    `json:"name,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Phone        string    `json:"Phone,omitempty" binding:"required" gorm:"type:varchar(20)"`
	Address      string    `json:"Address,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Contact      string    `json:"Contact,omitempty" binding:"required" gorm:"type:varchar(50)"`
	PhoneContact string    `json:"PhoneContact,omitempty" gorm:"type:varchar(50)"`
	MailContact  string    `json:"MailContact,omitempty" gorm:"type:varchar(256)"`
	NIT          string    `json:"NIT,omitempty" binding:"required" gorm:"type:varchar(20)"`
	Created      string    `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"Created"`
	Updated      string    `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"Updated"`
	DataState    DataState `json:"DataState" gorm:"foreignkey:data_state_id"`
	DataStateID  int       `json:-`
	City         City      `json:"City" gorm:"foreignkey:id"`
	CityID       int       `json:-`
}

//City data
type City struct {
	ID      int    `gorm:"primary_key" json:"ID,omitempty"`
	Name    string `json:"name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	State   State  `json:"State" gorm:"foreignkey:id"`
	StateID string `json:-`
}

//State data
type State struct {
	ID        int     `gorm:"primary_key" json:"ID,omitempty"`
	Name      string  `json:"name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Country   Country `json:"Country" gorm:"foreignkey:id"`
	CountryID string  `json:-`
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
	Name        string `json:"Dame,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(100)"`
}
