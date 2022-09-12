package votechain

import (
	"fmt"

	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func (r *Broker) Login(userLogin dto.Login) (*dto.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	users := sampledata.Users
	fmt.Println(userLogin)
	for _, user := range users {
		if user.Cuit == userLogin.Cuit && user.Password == userLogin.Password {
			fmt.Println("login success")
			return user, nil
		}
	}

	return nil, nil
}
