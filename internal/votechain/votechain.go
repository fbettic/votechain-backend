package votechain

import (
	"sync"

	"github.com/MatiasCermak/votechain-contracts-api/contracts/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

type Votechain interface {
	FetchOptions() ([]*dto.Option, error)
	Login(userLogin dto.Login) (*dto.User, error)
	RegisterVote(vote dto.Vote) (string, error)
	FetchOptionCount(*dto.Option) (*dto.OptionWithCount, error)
	GetVote(string) (*dto.Option, error)
}

type Broker struct {
	cfg     Config
	mtx     sync.RWMutex
	options map[string]*dto.Option
	conn    *api.Api
	client  *ethclient.Client
}

func New(cfg Config, options map[string]*dto.Option, local bool) (*Broker, error) {
	r := &Broker{}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	net := "http://votechain.ddns.net:8545"

	if local {
		net = "http://192.168.1.200:8545"
	}

	client, err := ethclient.Dial(net)

	if err != nil {
		panic(err)
	}

	//conn, err := api.NewApi(common.HexToAddress("0x00fFD3548725459255f1e78A61A07f1539Db0271"), client)
	conn, err := api.NewApi(common.HexToAddress("0x899CE22c2142f60ecA9574c3781A076136D46373"), client)
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

