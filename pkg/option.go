package option

type Option struct {
	ID          string `json:"ID"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
}

type OptionRepository interface {
	FetchOptions() ([]*Option, error)
}