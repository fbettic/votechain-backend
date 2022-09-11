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
	hasVoted, err := r.conn.HasVoted(&bind.CallOpts{Pending: false, From: fromAddress}, vote.Username)
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

	transaction, err := r.conn.CastVote(auth, vote.Username, vote.OptionID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Voto registrado correctamente")
	
	/*voted, err := r.conn.GetVoteCount(&bind.CallOpts{Pending: false, From: fromAddress}, vote.OptionID)
	if err != nil {
		panic(err)
	}

	fmt.Println(voted)*/

	return transaction, nil
}
