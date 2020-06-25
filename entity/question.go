package entity

//Questions user info
type Questions struct {
	ID          int16  `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(300)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(500)"`
}
