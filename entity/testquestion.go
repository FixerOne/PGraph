package entity

//Testquestion user info
type Testquestion struct {
	ID           int16      `gorm:"primary_key;auto_increment" json:"ID,omitempty"`
	Teststypes   Teststypes `json:"Teststype" gorm:"foreignkey:teststypes_id"`
	TeststypesID string     `json:"teststypes_id,omitempty"`
	Questions    Questions  `json:"Question" gorm:"foreignkey:questions_id"`
	QuestionID   string     `json:"questions_id,omitempty"`
	Relevant     bool       `json:"relevant"`
}
