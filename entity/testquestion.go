package entity

//Testquestion user info
type Testquestion struct {
	ID                   int16              `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Teststypes           Teststypes         `json:"Teststype" gorm:"foreignkey:teststypes_id"`
	TeststypesID         string             `json:"Teststypes_id,omitempty"`
	Basetestsquestions   Basetestsquestions `json:"Basetestsquestion" gorm:"foreignkey:basetestsquestions_id"`
	BasetestsquestionsID string             `json:"Basetestsquestions_id,omitempty"`
	Relevant             bool               `json:"Relevant"`
}
