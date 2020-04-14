package entity

//Project user info
type Project struct {
	ID          int16   `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Name        string  `json:"Name,omitempty" binding:"required" gorm:"type:varchar(100)"`
	Company     Company `json:"Company" gorm:"foreignkey:company_id"`
	CompanyID   string  `json:"company_id,omitempty"`
	InitDate    string  `json:"InitDate,omitempty" binding:"required" gorm:"type:date(100)"`
	FinishDate  string  `json:"FinishDate,omitempty" binding:"required" gorm:"type:date(100)"`
	ProjectType string  `json:"DocumentType,omitempty"`
	Created     string  `json:"Created" gorm:"default:CURRENT_TIMESTAMP"`
	Updated     string  `json:"Updated" gorm:"default:CURRENT_TIMESTAMP"`
}
