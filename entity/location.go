package entity

//City data
type City struct {
	ID      int    `gorm:"primary_key" json:"ID,omitempty"`
	Name    string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	State   State  `json:"State" gorm:"foreignkey:state_id"`
	StateID string `json:"StateID,omitempty"`
}

//State data
type State struct {
	ID        int     `gorm:"primary_key" json:"ID,omitempty"`
	Name      string  `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
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
