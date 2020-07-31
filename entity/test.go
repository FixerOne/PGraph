package entity

//Test company info
type Test struct {
	ID         int     `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Created    string  `json:"Created" gorm:"default:CURRENT_TIMESTAMP"`
	Updated    string  `json:"Updated" gorm:"default:CURRENT_TIMESTAMP"`
	UserID     string  `json:"user_id,omitempty" binding:"required" gorm:"type:int4"`
	User       User    `json:"User" gorm:"foreignkey:user_id"`
	ProjectID  string  `json:"project_id,omitempty" binding:"required" gorm:"type:int4"`
	Project    Project `json:"Project" gorm:"foreignkey:project_id"`
	InitHour   string  `json:"init_hour,omitempty" binding:"required" gorm:"type:int4"`
	FinishHour string  `json:"finish_hour,omitempty" binding:"required" gorm:"type:int4"`
	PDF        string  `json:"pdf,omitempty" binding:"required" gorm:"type:varchar(500)"`
	Word       string  `json:"word,omitempty" binding:"required" gorm:"type:varchar(500)"`
	Downloads  string  `json:"downloads,omitempty" binding:"required" gorm:"type:int4"`
	Date       string  `json:"date,omitempty" binding:"required" gorm:"type:date"`
}

/*id
user_id
project_id
date
init_hour
finish_hour
teststypes_id
created
updated
interviewer_user_id
job
record
pdf
word*/
