package votechain

import "github.com/fbettic/votechain-backend/pkg/dto"

func (r *Broker) FetchOptions() ([]*dto.Option, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	options := make([]*dto.Option, 0, len(r.options))
	for _, option := range r.options {
		options = append(options, option)
	}

	return options, nil
}
