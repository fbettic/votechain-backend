package webserver

import votechain "github.com/fbettic/votechain-backend/internal/votechain"

// Config for the storefront API.
type Config struct {
	Votechain votechain.Config
	Mock      bool

	Port int
}

func validateConfig(cfg Config) error {
	// This function would be used to validate a hydrated configuration; return an error if its invalid.
	return nil
}
