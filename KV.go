// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KVStorage

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
	_ = abi.ConvertType
)

// KVStorageMetaData contains all meta data concerning the KVStorage contract.
var KVStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506107b98061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063e942b51614610038578063fc2525ab14610054575b5f80fd5b610052600480360381019061004d9190610249565b610084565b005b61006e60048036038101906100699190610321565b6100f4565b60405161007b9190610408565b60405180910390f35b81815f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2086866040516100d2929190610464565b908152602001604051809103902091826100ed9291906106b6565b5050505050565b60605f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208383604051610142929190610464565b9081526020016040518091039020805461015b906104e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610187906104e0565b80156101d25780601f106101a9576101008083540402835291602001916101d2565b820191905f5260205f20905b8154815290600101906020018083116101b557829003601f168201915b505050505090509392505050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610209576102086101e8565b5b8235905067ffffffffffffffff811115610226576102256101ec565b5b602083019150836001820283011115610242576102416101f0565b5b9250929050565b5f805f8060408587031215610261576102606101e0565b5b5f85013567ffffffffffffffff81111561027e5761027d6101e4565b5b61028a878288016101f4565b9450945050602085013567ffffffffffffffff8111156102ad576102ac6101e4565b5b6102b9878288016101f4565b925092505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102f0826102c7565b9050919050565b610300816102e6565b811461030a575f80fd5b50565b5f8135905061031b816102f7565b92915050565b5f805f60408486031215610338576103376101e0565b5b5f6103458682870161030d565b935050602084013567ffffffffffffffff811115610366576103656101e4565b5b610372868287016101f4565b92509250509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156103b557808201518184015260208101905061039a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6103da8261037e565b6103e48185610388565b93506103f4818560208601610398565b6103fd816103c0565b840191505092915050565b5f6020820190508181035f83015261042081846103d0565b905092915050565b5f81905092915050565b828183375f83830152505050565b5f61044b8385610428565b9350610458838584610432565b82840190509392505050565b5f610470828486610440565b91508190509392505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806104f757607f821691505b60208210810361050a576105096104b3565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261056c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610531565b6105768683610531565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6105ba6105b56105b08461058e565b610597565b61058e565b9050919050565b5f819050919050565b6105d3836105a0565b6105e76105df826105c1565b84845461053d565b825550505050565b5f90565b6105fb6105ef565b6106068184846105ca565b505050565b5b818110156106295761061e5f826105f3565b60018101905061060c565b5050565b601f82111561066e5761063f81610510565b61064884610522565b81016020851015610657578190505b61066b61066385610522565b83018261060b565b50505b505050565b5f82821c905092915050565b5f61068e5f1984600802610673565b1980831691505092915050565b5f6106a6838361067f565b9150826002028217905092915050565b6106c0838361047c565b67ffffffffffffffff8111156106d9576106d8610486565b5b6106e382546104e0565b6106ee82828561062d565b5f601f83116001811461071b575f8415610709578287013590505b610713858261069b565b86555061077a565b601f19841661072986610510565b5f5b828110156107505784890135825560018201915060208501945060208101905061072b565b8683101561076d5784890135610769601f89168261067f565b8355505b6001600288020188555050505b5050505050505056fea26469706673582212205bc2ac49974afbc8fcfe617adf1b83fc02f5c412fe65f782ca6544ca863900db64736f6c63430008180033",
}

// KVStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use KVStorageMetaData.ABI instead.
var KVStorageABI = KVStorageMetaData.ABI

// KVStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KVStorageMetaData.Bin instead.
var KVStorageBin = KVStorageMetaData.Bin

// DeployKVStorage deploys a new Ethereum contract, binding an instance of KVStorage to it.
func DeployKVStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KVStorage, error) {
	parsed, err := KVStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KVStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KVStorage{KVStorageCaller: KVStorageCaller{contract: contract}, KVStorageTransactor: KVStorageTransactor{contract: contract}, KVStorageFilterer: KVStorageFilterer{contract: contract}}, nil
}

// KVStorage is an auto generated Go binding around an Ethereum contract.
type KVStorage struct {
	KVStorageCaller     // Read-only binding to the contract
	KVStorageTransactor // Write-only binding to the contract
	KVStorageFilterer   // Log filterer for contract events
}

// KVStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KVStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KVStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KVStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KVStorageSession struct {
	Contract     *KVStorage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KVStorageCallerSession struct {
	Contract *KVStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KVStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KVStorageTransactorSession struct {
	Contract     *KVStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KVStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KVStorageRaw struct {
	Contract *KVStorage // Generic contract binding to access the raw methods on
}

// KVStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KVStorageCallerRaw struct {
	Contract *KVStorageCaller // Generic read-only contract binding to access the raw methods on
}

// KVStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KVStorageTransactorRaw struct {
	Contract *KVStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVStorage creates a new instance of KVStorage, bound to a specific deployed contract.
func NewKVStorage(address common.Address, backend bind.ContractBackend) (*KVStorage, error) {
	contract, err := bindKVStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVStorage{KVStorageCaller: KVStorageCaller{contract: contract}, KVStorageTransactor: KVStorageTransactor{contract: contract}, KVStorageFilterer: KVStorageFilterer{contract: contract}}, nil
}

// NewKVStorageCaller creates a new read-only instance of KVStorage, bound to a specific deployed contract.
func NewKVStorageCaller(address common.Address, caller bind.ContractCaller) (*KVStorageCaller, error) {
	contract, err := bindKVStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVStorageCaller{contract: contract}, nil
}

// NewKVStorageTransactor creates a new write-only instance of KVStorage, bound to a specific deployed contract.
func NewKVStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*KVStorageTransactor, error) {
	contract, err := bindKVStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVStorageTransactor{contract: contract}, nil
}

// NewKVStorageFilterer creates a new log filterer instance of KVStorage, bound to a specific deployed contract.
func NewKVStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*KVStorageFilterer, error) {
	contract, err := bindKVStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVStorageFilterer{contract: contract}, nil
}

// bindKVStorage binds a generic wrapper to an already deployed contract.
func bindKVStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KVStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVStorage *KVStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KVStorage.Contract.KVStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVStorage *KVStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVStorage.Contract.KVStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVStorage *KVStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVStorage.Contract.KVStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVStorage *KVStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KVStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVStorage *KVStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVStorage *KVStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVStorage.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xfc2525ab.
//
// Solidity: function get(address origin, string key) view returns(string)
func (_KVStorage *KVStorageCaller) Get(opts *bind.CallOpts, origin common.Address, key string) (string, error) {
	var out []interface{}
	err := _KVStorage.contract.Call(opts, &out, "get", origin, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xfc2525ab.
//
// Solidity: function get(address origin, string key) view returns(string)
func (_KVStorage *KVStorageSession) Get(origin common.Address, key string) (string, error) {
	return _KVStorage.Contract.Get(&_KVStorage.CallOpts, origin, key)
}

// Get is a free data retrieval call binding the contract method 0xfc2525ab.
//
// Solidity: function get(address origin, string key) view returns(string)
func (_KVStorage *KVStorageCallerSession) Get(origin common.Address, key string) (string, error) {
	return _KVStorage.Contract.Get(&_KVStorage.CallOpts, origin, key)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_KVStorage *KVStorageTransactor) Set(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _KVStorage.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_KVStorage *KVStorageSession) Set(key string, value string) (*types.Transaction, error) {
	return _KVStorage.Contract.Set(&_KVStorage.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_KVStorage *KVStorageTransactorSession) Set(key string, value string) (*types.Transaction, error) {
	return _KVStorage.Contract.Set(&_KVStorage.TransactOpts, key, value)
}
