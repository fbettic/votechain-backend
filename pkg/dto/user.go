package dto

type Login struct {
	Cuit 				int `json:"cuit"`
	Password 		string `json:"password"`
}

type User struct{
	Name				string `json:"name"`
	Cuit 				int `json:"cuit"`
	Password 		string `json:"password,omitempty"`
	AccessToken string `json:"access_token"`
	HasVoted	bool	`json:"has_voted,omitemty"`
}