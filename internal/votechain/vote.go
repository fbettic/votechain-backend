package votechain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbettic/votechain-backend/internal/verification"
	"github.com/fbettic/votechain-backend/pkg/dto"
	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
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
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("Invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := r.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := r.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainid, _ := r.client.ChainID(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	isvalid, err := r.conn.IsValidOption(&bind.CallOpts{Pending: false, From: fromAddress},vote.OptionID)
	if err != nil {
		panic(err)
	}
	if !isvalid{
		return "", errors.New("400 - Invalid option selected")
	}
	hasVoted, err := r.conn.HasVoted(&bind.CallOpts{Pending: false, From: fromAddress}, hash)
	if err != nil {
		panic(err)
	}
	if hasVoted{
		return "", errors.New("401 - User has alredy voted")
	}

	_, err = r.conn.CastVote(auth, hash, vote.OptionID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Voto registrado correctamente")

	validationCode, err := verification.CreateVerificationCode(hash)
	if err != nil {
		return "", err
	}

	return validationCode, nil
}

func (r *Broker) GetVote(code string) (*dto.Option, error){

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
		if opt.ID == optionVoted{
			option = opt
		}
	}

	return option, nil
}