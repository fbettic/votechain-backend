package dto

type Answer struct {
	QuestionID		string		`json:"question:_id"`
	Answer			string		`json:"answer"`
}

type Answers struct {
	Answers 		[]Answer	`json:"answers"`
	AccessToken 	string 		`json:"access_token"`
}