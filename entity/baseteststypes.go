package entity

//Baseteststypes user info
type Baseteststypes struct {
	ID          int16  `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(500)"`
}
