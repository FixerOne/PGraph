package entity

//Teststypes user info
type Teststypes struct {
	ID          int16  `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"Description,omitempty" binding:"required" gorm:"type:varchar(500)"`
	Title       string `json:"Title,omitempty" binding:"required" gorm:"type:varchar(500)"`
}
