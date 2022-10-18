// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"num\",\"type\":\"string[]\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StorageBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StorageBinRuntime = `608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b578063ae00106a14610059575b600080fd5b61004361006e565b6040516100509190610223565b60405180910390f35b61006c610067366004610305565b610147565b005b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561013e5783829060005260206000200180546100b19061041e565b80601f01602080910402602001604051908101604052809291908181526020018280546100dd9061041e565b801561012a5780601f106100ff5761010080835404028352916020019161012a565b820191906000526020600020905b81548152906001019060200180831161010d57829003601f168201915b505050505081526020019060010190610092565b50505050905090565b805161015a90600090602084019061015e565b5050565b8280548282559060005260206000209081019282156101a4579160200282015b828111156101a4578251829061019490826104a7565b509160200191906001019061017e565b506101b09291506101b4565b5090565b808211156101b05760006101c882826101d1565b506001016101b4565b5080546101dd9061041e565b6000825580601f106101ed575050565b601f01602090049060005260206000209081019061020b919061020e565b50565b5b808211156101b0576000815560010161020f565b6000602080830181845280855180835260408601915060408160051b87010192508387016000805b838110156102b057888603603f1901855282518051808852835b81811015610280578281018a01518982018b01528901610265565b8181111561029057848a838b0101525b50601f01601f19169690960187019550938601939186019160010161024b565b509398975050505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156102fd576102fd6102be565b604052919050565b6000602080838503121561031857600080fd5b823567ffffffffffffffff8082111561033057600080fd5b8185019150601f868184011261034557600080fd5b823582811115610357576103576102be565b8060051b6103668682016102d4565b918252848101860191868101908a84111561038057600080fd5b87870192505b838310156104105782358681111561039e5760008081fd5b8701603f81018c136103b05760008081fd5b888101356040888211156103c6576103c66102be565b6103d7828901601f19168c016102d4565b8281528e828486010111156103ec5760008081fd5b828285018d83013760009281018c0192909252508352509187019190870190610386565b9a9950505050505050505050565b600181811c9082168061043257607f821691505b60208210810361045257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156104a257600081815260208120601f850160051c8101602086101561047f5750805b601f850160051c820191505b8181101561049e5782815560010161048b565b5050505b505050565b815167ffffffffffffffff8111156104c1576104c16102be565b6104d5816104cf845461041e565b84610458565b602080601f83116001811461050a57600084156104f25750858301515b600019600386901b1c1916600185901b17855561049e565b600085815260208120601f198616915b828110156105395788860151825594840194600190910190840161051a565b50858210156105575787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220810577e9bbcd780f37a7a0c90d1851df8857ea39888052ef668e28e2727f906164736f6c634300080f0033`

// StorageFuncSigs maps the 4-byte function signature to its string representation.
var StorageFuncSigs = map[string]string{
	"2e64cec1": "retrieve()",
	"ae00106a": "store(string[])",
}

// StorageBin is the compiled bytecode used for deploying new contracts.
var StorageBin = "0x608060405234801561001057600080fd5b5061059d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b578063ae00106a14610059575b600080fd5b61004361006e565b6040516100509190610223565b60405180910390f35b61006c610067366004610305565b610147565b005b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561013e5783829060005260206000200180546100b19061041e565b80601f01602080910402602001604051908101604052809291908181526020018280546100dd9061041e565b801561012a5780601f106100ff5761010080835404028352916020019161012a565b820191906000526020600020905b81548152906001019060200180831161010d57829003601f168201915b505050505081526020019060010190610092565b50505050905090565b805161015a90600090602084019061015e565b5050565b8280548282559060005260206000209081019282156101a4579160200282015b828111156101a4578251829061019490826104a7565b509160200191906001019061017e565b506101b09291506101b4565b5090565b808211156101b05760006101c882826101d1565b506001016101b4565b5080546101dd9061041e565b6000825580601f106101ed575050565b601f01602090049060005260206000209081019061020b919061020e565b50565b5b808211156101b0576000815560010161020f565b6000602080830181845280855180835260408601915060408160051b87010192508387016000805b838110156102b057888603603f1901855282518051808852835b81811015610280578281018a01518982018b01528901610265565b8181111561029057848a838b0101525b50601f01601f19169690960187019550938601939186019160010161024b565b509398975050505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156102fd576102fd6102be565b604052919050565b6000602080838503121561031857600080fd5b823567ffffffffffffffff8082111561033057600080fd5b8185019150601f868184011261034557600080fd5b823582811115610357576103576102be565b8060051b6103668682016102d4565b918252848101860191868101908a84111561038057600080fd5b87870192505b838310156104105782358681111561039e5760008081fd5b8701603f81018c136103b05760008081fd5b888101356040888211156103c6576103c66102be565b6103d7828901601f19168c016102d4565b8281528e828486010111156103ec5760008081fd5b828285018d83013760009281018c0192909252508352509187019190870190610386565b9a9950505050505050505050565b600181811c9082168061043257607f821691505b60208210810361045257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156104a257600081815260208120601f850160051c8101602086101561047f5750805b601f850160051c820191505b8181101561049e5782815560010161048b565b5050505b505050565b815167ffffffffffffffff8111156104c1576104c16102be565b6104d5816104cf845461041e565b84610458565b602080601f83116001811461050a57600084156104f25750858301515b600019600386901b1c1916600185901b17855561049e565b600085815260208120601f198616915b828110156105395788860151825594840194600190910190840161051a565b50858210156105575787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220810577e9bbcd780f37a7a0c90d1851df8857ea39888052ef668e28e2727f906164736f6c634300080f0033"

// DeployStorage deploys a new Klaytn contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around a Klaytn contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around a Klaytn contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around a Klaytn contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Storage *StorageCaller) Retrieve(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "retrieve")
	return *ret0, err
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Storage *StorageSession) Retrieve() ([]string, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Storage *StorageCallerSession) Retrieve() ([]string, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] num) returns()
func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, num []string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] num) returns()
func (_Storage *StorageSession) Store(num []string) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] num) returns()
func (_Storage *StorageTransactorSession) Store(num []string) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, num)
}
