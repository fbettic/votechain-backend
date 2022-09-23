package votechain

import (
	"errors"
	"fmt"
	"strconv"

	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
	"github.com/fbettic/votechain-backend/internal/verification"
	"github.com/fbettic/votechain-backend/pkg/dto"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func (r *Broker) Login(userLogin dto.Login) (*dto.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	users := sampledata.Users
	for _, user := range users {
		if user.Cuit == userLogin.Cuit && user.Password == userLogin.Password {
			return user, nil
		}
	}

	return nil, errors.New("404 - User not found")
}

func (r *Broker) Register(newUser dto.User) error {
	users := sampledata.Users
	for _, user := range users {
		if newUser.Cuit == user.Cuit {
			return errors.New("409 - Cuit already registered")
		}
	}

	err := passwordvalidator.Validate(newUser.Password, 60)
	if err != nil {
		fmt.Println(err)
		return err
	}

	token, err := verification.GenerateJWT(newUser.Name, newUser.Cuit)
	if err != nil {
		return err
	}

	newUser.AccessToken = token

	users[strconv.Itoa(len(users)+1)] = &newUser

	return nil
}
