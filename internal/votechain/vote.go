package votechain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func (r *Broker) RegisterVote(vote dto.Vote) (*types.Transaction, *dto.ErrorMessage) {
	/*hash, err := GetHashCode(code)
	if err != nil {
		msgerr := &dto.ErrorMessage{
					Status: 404,
					Message: "Code not found",
				}
		return nil, msgerr
	}
	*/
	hash := "576as5d"

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
		erro := &dto.ErrorMessage{
					Status: 400,
					Message: "Invalid option selected",
				}
		return nil, erro
	}
	hasVoted, err := r.conn.HasVoted(&bind.CallOpts{Pending: false, From: fromAddress}, hash)
	if err != nil {
		panic(err)
	}
	if hasVoted{
		erro := &dto.ErrorMessage{
					Status: 401,
					Message: "User has alredy voted",
				}
		return nil, erro
	}

	transaction, err := r.conn.CastVote(auth, hash, vote.OptionID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Voto registrado correctamente")

	return transaction, nil
}

func (r *Broker) GetVote(code string) (*dto.Option, *dto.ErrorMessage){

	/*hash, err := GetHashCode(code)
	if err != nil {
		msgerr := &dto.ErrorMessage{
					Status: 404,
					Message: "Code not found",
				}
		return nil, msgerr
	}
	*/
	hash := "576as5d"
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
		msgerr := &dto.ErrorMessage{
					Status: 400,
					Message: "User has not voted yet",
				}
		return nil, msgerr
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