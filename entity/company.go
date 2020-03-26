package entity

//Company company info
type Company struct {
	ID           int    `gorm:"primary_key;auto_increment" json:"_id,omitempty"`
	Name         string `json:"name,omitempty" binding:"required"`
	Phone        string `json:"Phone,omitempty" binding:"required"`
	Address      string `json:"Address,omitempty" binding:"required"`
	CityID       string `json:"CityID,omitempty" binding:"required"`
	Contact      string `json:"Contact,omitempty" binding:"required"`
	PhoneContact string `json:"PhoneContact,omitempty"`
	MailContact  string `json:"MailContact,omitempty"`
	NIT          string `json:"NIT,omitempty" binding:"required"`
	Created      string `json:"Created,omitempty"`
	Updated      string `json:"Updated,omitempty"`
	State        string `json:"State,omitempty" binding:"required"`
}
