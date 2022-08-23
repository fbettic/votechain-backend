package inmem

import (
	"context"
	"crypto/ecdsa"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	option "github.com/fbettic/votechain-backend/pkg"
	"github.com/fbettic/votechain-backend/pkg/api"
)

type optionRepository struct {
	mtx     sync.RWMutex
	options map[string]*option.Option
	conn    *api.Api
	client  *ethclient.Client
}

func NewOptionRepository(options map[string]*option.Option) option.OptionRepository {

	client, err := ethclient.Dial("http://votechain.ddns.net:8545")

	if err != nil {
		panic(err)
	}

	conn, err := api.NewApi(common.HexToAddress("0x00fFD3548725459255f1e78A61A07f1539Db0271"), client)
	if err != nil {
		panic(err)
	}

	if options == nil {
		options = make(map[string]*option.Option)
	}

	return &optionRepository{
		options: options,
		conn:    conn,
		client:  client,
	}
}

func (r *optionRepository) FetchOptions() ([]*option.Option, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	options := make([]*option.Option, 0, len(r.options))
	for _, option := range r.options {
		options = append(options, option)
	}

	return options, nil
}

func (r *optionRepository) RegisterVote(vote option.Vote) *types.Transaction {
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
	fun := func(common common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		chainid, _ := r.client.ChainID(context.Background())
		signer := types.NewEIP155Signer(chainid)
		transaction.WithSignature(signer)
		return transaction, nil
	}

	transaction, err := r.conn.CastVote(&bind.TransactOpts{From: fromAddress, Signer: fun}, vote.Username, vote.OptionID)
	if err != nil {
		panic(err)
	}

	return transaction
}
