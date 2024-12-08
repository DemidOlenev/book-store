package dto


type ClientDto struct {
	Id int `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Phone string `json:"phone"`
}
