package dto

type Option struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
}
