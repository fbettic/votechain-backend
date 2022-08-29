package votechain

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbettic/votechain-backend/pkg/api"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

type Votechain interface {
	FetchOptions() ([]*dto.Option, error)
	RegisterVote(vote dto.Vote) *types.Transaction
}

type Broker struct {
	cfg     Config
	mtx     sync.RWMutex
	options map[string]*dto.Option
	conn    *api.Api
	client  *ethclient.Client
}

func New(cfg Config, options map[string]*dto.Option) (*Broker, error) {
	r := &Broker{}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	client, err := ethclient.Dial("http://votechain.ddns.net:8545")

	if err != nil {
		panic(err)
	}

	conn, err := api.NewApi(common.HexToAddress("0x00fFD3548725459255f1e78A61A07f1539Db0271"), client)
	if err != nil {
		panic(err)
	}

	if options == nil {
		options = make(map[string]*dto.Option)
	}

	r.client = client
	r.conn = conn
	r.options = options
	r.cfg = cfg

	return r, nil
}
