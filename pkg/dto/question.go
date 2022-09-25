package dto

type QuestionType string

const (
	Undefined	QuestionType = ""
	Choice					 = "Choice"
	Written					 = "Written"
)

type ChoiceQuestion struct {
	ID				string			`json:"id"`
	Question 		string			`json:"question"`
	Choices 		[]string		`json:"choices,omitempty"`
	Type 			QuestionType	`json:"type"`
}
