package dto

type Proposal struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type Option struct {
	ID           string     `json:"id"`
	Name         string     `json:"name,omitempty"`
	Image        string     `json:"image,omitempty"`
	Group        string     `json:"group,omitempty"`
	Presentation string     `json:"presentation,omitempty"`
	Description  string     `json:"description,omitempty"`
	Proposals    []Proposal `json:"proposals"`
}

type OptionWithCount struct {
	ID           string `json:"id"`
	Name         string `json:"name,omitempty"`
	Image        string `json:"image,omitempty"`
	Group        string `json:"group,omitempty"`
	Presentation string `json:"presentation,omitempty"`
	Description  string `json:"description,omitempty"`
	Votes        string `json:"votes"`
}
