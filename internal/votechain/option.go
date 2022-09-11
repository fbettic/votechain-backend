package votechain

import (
	"crypto/ecdsa"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func (r *Broker) FetchOptions() ([]*dto.Option, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	privateKey, err := crypto.HexToECDSA("8bbbb1b345af56b560a5b20bd4b0ed1cd8cc9958a16262bc75118453cb546df7")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainOpt, err := r.conn.GetOptions(&bind.CallOpts{Pending: false, From: fromAddress})
	if err != nil {
		panic(err)
	}

	options := make([]*dto.Option, 0, len(r.options))
	for _, option := range r.options {
		for i := range chainOpt {
			if chainOpt[i] == option.ID {
				options = append(options, option)
			}
		}
	}

	return options, nil
}

func (r *Broker) FetchOptionCount(option *dto.Option) (*dto.OptionWithCount, *dto.ErrorMessage) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	privateKey, err := crypto.HexToECDSA("8bbbb1b345af56b560a5b20bd4b0ed1cd8cc9958a16262bc75118453cb546df7")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	isValid, err := r.conn.IsValidOption(&bind.CallOpts{Pending: false, From: fromAddress}, option.ID)
	if err != nil {
		panic(err)
	}
	if !isValid {
		erro := &dto.ErrorMessage{
			Status:  400,
			Message: "Invalid option selected",
		}
		return nil, erro
	}
	count, err := r.conn.GetVoteCount(&bind.CallOpts{Pending: false, From: fromAddress}, option.ID)
	if err != nil {
		panic(err)
	}
	optionWithCount := new(dto.OptionWithCount)
	optionWithCount.Description = option.Description
	optionWithCount.Group = option.Group
	optionWithCount.ID = option.ID
	optionWithCount.Image = option.Image
	optionWithCount.Name = option.Name
	optionWithCount.Votes = strconv.FormatUint(uint64(count), 10)

	return optionWithCount, nil
}
