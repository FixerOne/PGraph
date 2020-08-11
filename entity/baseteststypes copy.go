package entity

//Userstypes user info
type Userstypes struct {
	ID          int16  `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(50)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(100)"`
}
