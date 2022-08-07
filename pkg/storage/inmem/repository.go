package inmem

import (
	"fmt"
	"sync"

	option "github.com/fbettic/votechain-backend/pkg"
)

type optionRepository struct {
	mtx    sync.RWMutex
	options map[string]*option.Option
}

func NewOptionRepository(options map[string]*option.Option) option.OptionRepository {
	if options == nil {
		options = make(map[string]*option.Option)
	}
	
	return &optionRepository{
		options: options,
	}
}

func (r *optionRepository) FetchOptions() ([]*option.Option, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	options := make([]*option.Option, 0, len(r.options))
	for _, option := range r.options {
		options = append(options, option)
	}

	return options, nil
}

func (r *optionRepository) RegisterVote(vote option.Vote) {
	fmt.Print(vote)
}