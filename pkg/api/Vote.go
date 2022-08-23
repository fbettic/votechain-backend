// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_options\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e7438038062000e748339810160408190526200003491620001cc565b6200003f336200005c565b805162000054906003906020840190620000ac565b505062000473565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054828255906000526020600020908101928215620000f7579160200282015b82811115620000f75782518290620000e69082620003a7565b5091602001919060010190620000cd565b506200010592915062000109565b5090565b80821115620001055760006200012082826200012a565b5060010162000109565b508054620001389062000318565b6000825580601f1062000149575050565b601f0160209004906000526020600020908101906200016991906200016c565b50565b5b808211156200010557600081556001016200016d565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620001c457620001c462000183565b604052919050565b60006020808385031215620001e057600080fd5b82516001600160401b0380821115620001f857600080fd5b8185019150601f86818401126200020e57600080fd5b82518281111562000223576200022362000183565b8060051b6200023486820162000199565b918252848101860191868101908a8411156200024f57600080fd5b87870192505b838310156200030a578251868111156200026f5760008081fd5b8701603f81018c13620002825760008081fd5b888101518781111562000299576200029962000183565b620002ac818801601f19168b0162000199565b81815260408e81848601011115620002c45760008081fd5b60005b83811015620002e4578481018201518382018e01528c01620002c7565b83811115620002f65760008d85850101525b505084525050918701919087019062000255565b9a9950505050505050505050565b600181811c908216806200032d57607f821691505b6020821081036200034e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620003a257600081815260208120601f850160051c810160208610156200037d5750805b601f850160051c820191505b818110156200039e5782815560010162000389565b5050505b505050565b81516001600160401b03811115620003c357620003c362000183565b620003db81620003d4845462000318565b8462000354565b602080601f831160018114620004135760008415620003fa5750858301515b600019600386901b1c1916600185901b1785556200039e565b600085815260208120601f198616915b82811015620004445788860151825594840194600190910190840162000423565b5085821015620004635787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6109f180620004836000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063069953a714610067578063715018a6146100945780638da5cb5b1461009e578063baa40e5c146100b9578063c479fc91146100d9578063f2fde38b146100ec575b600080fd5b61007a6100753660046105fd565b6100ff565b60405163ffffffff90911681526020015b60405180910390f35b61009c610164565b005b6000546040516001600160a01b03909116815260200161008b565b6100cc6100c73660046105fd565b610178565b60405161008b919061066a565b61009c6100e736600461069d565b610285565b61009c6100fa366004610701565b61038a565b6000610109610403565b6101128261045d565b6101375760405162461bcd60e51b815260040161012e90610731565b60405180910390fd5b6002826040516101479190610772565b9081526040519081900360200190205463ffffffff169050919050565b61016c610403565b61017660006104d4565b565b6060610182610403565b61018b82610524565b6101d75760405162461bcd60e51b815260206004820152601960248201527f566f74653a20566f746572436f6465206e6f7420666f756e6400000000000000604482015260640161012e565b6001826040516101e79190610772565b908152602001604051809103902080546102009061078e565b80601f016020809104026020016040519081016040528092919081815260200182805461022c9061078e565b80156102795780601f1061024e57610100808354040283529160200191610279565b820191906000526020600020905b81548152906001019060200180831161025c57829003601f168201915b50505050509050919050565b61028d610403565b6102968161045d565b6102b25760405162461bcd60e51b815260040161012e90610731565b6102bb82610524565b156103085760405162461bcd60e51b815260206004820152601d60248201527f566f74653a20566f7465722068617320616c726561647920766f746564000000604482015260640161012e565b806001836040516103199190610772565b908152602001604051809103902090816103339190610817565b506002816040516103449190610772565b908152604051908190036020019020805463ffffffff16906000610367836108ed565b91906101000a81548163ffffffff021916908363ffffffff160217905550505050565b610392610403565b6001600160a01b0381166103f75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161012e565b610400816104d4565b50565b6000546001600160a01b031633146101765760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161012e565b6000805b60035460ff821610156104cb57828051906020012060038260ff168154811061048c5761048c610910565b906000526020600020016040516104a39190610926565b6040518091039020036104b95750600192915050565b806104c38161099c565b915050610461565b50600092915050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806001836040516105379190610772565b908152602001604051809103902080546105509061078e565b9050119050919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261058157600080fd5b813567ffffffffffffffff8082111561059c5761059c61055a565b604051601f8301601f19908116603f011681019082821181831017156105c4576105c461055a565b816040528381528660208588010111156105dd57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561060f57600080fd5b813567ffffffffffffffff81111561062657600080fd5b61063284828501610570565b949350505050565b60005b8381101561065557818101518382015260200161063d565b83811115610664576000848401525b50505050565b602081526000825180602084015261068981604085016020870161063a565b601f01601f19169190910160400192915050565b600080604083850312156106b057600080fd5b823567ffffffffffffffff808211156106c857600080fd5b6106d486838701610570565b935060208501359150808211156106ea57600080fd5b506106f785828601610570565b9150509250929050565b60006020828403121561071357600080fd5b81356001600160a01b038116811461072a57600080fd5b9392505050565b60208082526021908201527f566f74653a2056616c75652073656e74206973206e6f7420616e204f7074696f6040820152603760f91b606082015260800190565b6000825161078481846020870161063a565b9190910192915050565b600181811c908216806107a257607f821691505b6020821081036107c257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561081257600081815260208120601f850160051c810160208610156107ef5750805b601f850160051c820191505b8181101561080e578281556001016107fb565b5050505b505050565b815167ffffffffffffffff8111156108315761083161055a565b6108458161083f845461078e565b846107c8565b602080601f83116001811461087a57600084156108625750858301515b600019600386901b1c1916600185901b17855561080e565b600085815260208120601f198616915b828110156108a95788860151825594840194600190910190840161088a565b50858210156108c75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff808316818103610906576109066108d7565b6001019392505050565b634e487b7160e01b600052603260045260246000fd5b60008083546109348161078e565b6001828116801561094c576001811461096157610990565b60ff1984168752821515830287019450610990565b8760005260208060002060005b858110156109875781548a82015290840190820161096e565b50505082870194505b50929695505050505050565b600060ff821660ff81036109b2576109b26108d7565b6001019291505056fea26469706673582212200a375e9bd6eddab1443ad3b8bd8313bffa2b4f1a0ef7cc850003030766b8b70464736f6c634300080f0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _options []string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _options)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiCaller) GetVote(opts *bind.CallOpts, voterCode string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getVote", voterCode)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiSession) GetVote(voterCode string) (string, error) {
	return _Api.Contract.GetVote(&_Api.CallOpts, voterCode)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiCallerSession) GetVote(voterCode string) (string, error) {
	return _Api.Contract.GetVote(&_Api.CallOpts, voterCode)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiCaller) GetVoteCount(opts *bind.CallOpts, option string) (uint32, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getVoteCount", option)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiSession) GetVoteCount(option string) (uint32, error) {
	return _Api.Contract.GetVoteCount(&_Api.CallOpts, option)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiCallerSession) GetVoteCount(option string) (uint32, error) {
	return _Api.Contract.GetVoteCount(&_Api.CallOpts, option)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiTransactor) CastVote(opts *bind.TransactOpts, voterCode string, option string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "castVote", voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiSession) CastVote(voterCode string, option string) (*types.Transaction, error) {
	return _Api.Contract.CastVote(&_Api.TransactOpts, voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiTransactorSession) CastVote(voterCode string, option string) (*types.Transaction, error) {
	return _Api.Contract.CastVote(&_Api.TransactOpts, voterCode, option)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
