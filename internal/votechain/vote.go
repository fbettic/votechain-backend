package votechain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
	"github.com/fbettic/votechain-backend/internal/verification"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func (r *Broker) RegisterVote(vote dto.Vote) (string, error) {
	var user *dto.User
	user = nil
	for _,userData := range sampledata.Users{
		if userData.AccessToken == vote.AccessToken{
			user = userData
			break
		}
	}
	if user == nil{
		return "", errors.New("401 - Invalid login token")
	}

	hash, err := verification.CreateHash(*user)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA("8bbbb1b345af56b560a5b20bd4b0ed1cd8cc9958a16262bc75118453cb546df7")
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("500 - Invalid member node key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := r.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := r.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	chainid, _ := r.client.ChainID(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	option := r.options[vote.OptionID]
	if option == nil {
		return "", errors.New("404 - Option not found")
	}
	optionHash, err := verification.CreateHash(*RemoveImage(option))
	if err != nil {
		return "", err
	}

	isvalid, err := r.conn.IsValidOption(&bind.CallOpts{Pending: false, From: fromAddress},optionHash)
	if err != nil {
		return "", err
	}
	if !isvalid{
		return "", errors.New("400 - Invalid option selected")
	}
	hasVoted, err := r.conn.HasVoted(&bind.CallOpts{Pending: false, From: fromAddress}, hash)
	if err != nil {
		return "", err
	}
	if hasVoted{
		user.HasVoted = true
		return "", errors.New("401 - User has alredy voted")
	}

	_, err = r.conn.CastVote(auth, hash, optionHash)
	if err != nil {
		return "", err
	}

	fmt.Println("Voto registrado correctamente")

	validationCode, err := verification.CreateVerificationCode(hash)
	if err != nil {
		return "", err
	}

	user.HasVoted = true

	return validationCode, nil
}

func (r *Broker) GetVote(code string) (*dto.Option, error){

	fmt.Println(code)

	hash, err := verification.GetHashCode(code)
	if err != nil {
		return nil, err
	}

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
	
	hasvoted, err := r.conn.HasVoted(&bind.CallOpts{Pending: false, From: fromAddress}, hash)
	if err != nil{
		panic(err)
	}
	if !hasvoted {
		return nil, errors.New("400 - User has not voted yet")
	}

	optionVoted, err := r.conn.GetVote(&bind.CallOpts{Pending: false, From: fromAddress}, hash)
	if err != nil {
		panic(err)
	}

	option := new(dto.Option)
	for _, opt := range r.options{
		hash, err := verification.CreateHash(*RemoveImage(opt))
		if err != nil {
			return nil, err
		}
		if hash == optionVoted{
			option = opt
		}
	}

	return option, nil
}