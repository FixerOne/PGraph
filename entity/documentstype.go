package entity

//Documentstypes user info
type Documentstypes struct {
	ID           int16  `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name         string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Description  string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(500)"`
	Abbreviation string `json:"Abbreviation,omitempty" binding:"required" gorm:"type:varchar(3)"`
	Active       bool   `json:"Active"`
}
