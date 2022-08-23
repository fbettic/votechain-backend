package option

import "github.com/ethereum/go-ethereum/core/types"

type Option struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
}

type Vote struct {
	OptionID string `json:"option_id"`
	Username string `json:"username"`
}

type OptionRepository interface {
	FetchOptions() ([]*Option, error)
	RegisterVote(vote Vote) *types.Transaction
}
