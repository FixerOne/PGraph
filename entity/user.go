package entity

//User user info
type User struct {
	ID                   int16  `json:"_id,omitempty"`
	FirstName            string `json:"FirstName,omitempty"`
	LastName             string `json:"LastName,omitempty"`
	Birthday             string `json:"Birthday,omitempty"`
	Document             string `json:"Document,omitempty"`
	DocumentType         string `json:"DocumentType,omitempty"`
	City                 string `json:"City,omitempty"`
	Address              string `json:"Address,omitempty"`
	NationalityCountryID string `json:"Nationality_Country_Id,omitempty"`
	Gender               string `json:"Gender,omitempty"`
	Company              string `json:"Company,omitempty"`
	Phone                string `json:"Phone,omitempty"`
	Email                string `json:"Email,omitempty"`
	UserType             string `json:"UserType,omitempty"`
	Created              string `json:"Created,omitempty"`
	Updated              string `json:"Updated,omitempty"`
	State                string `json:"State,omitempty"`
	Password             string `json:"Password,omitempty"`
}

/*//UserResponseMessage users list info
type UserResponseMessage struct {
	Message ResponseMessage
	User    User
}

//UsersResponseMessage users list info
type UsersResponseMessage struct {
	Message ResponseMessage
	Users   []User
}

//ResponseMessage message
type ResponseMessage struct {
	Error      string `json:"Error"`
	Message    string `json:"Message"`
	MethodName string `json:"MethodName"`
	Count      int    `json:"Count"`
}*/
