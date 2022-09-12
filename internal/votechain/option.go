package votechain

import (
	"fmt"
	"strconv"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/fbettic/votechain-backend/pkg/dto"
	"github.com/ethereum/go-ethereum/crypto"
)

func (r *Broker) FetchOptions() ([]*dto.Option, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	options := make([]*dto.Option, 0, len(r.options))
	for _, option := range r.options {
		fmt.Println(option)
		options = append(options, option)
	}

	

	return options, nil
}

func (r *Broker) FetchOptionCount(option *dto.Option) (*dto.OptionWithCount, error){
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
		fmt.Println("Option invalid")
		return nil, fmt.Errorf("Option invalid")
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