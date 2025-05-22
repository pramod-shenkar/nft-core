// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dlt

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

// RequestmodelRequest is an auto generated low-level Go binding around an user-defined struct.
type RequestmodelRequest struct {
	StorageId *big.Int
	Hashdata  string
	Creator   common.Address
	Status    uint8
}

// TokenmodelToken is an auto generated low-level Go binding around an user-defined struct.
type TokenmodelToken struct {
	StorageId *big.Int
	Hashdata  string
	Owner     common.Address
	Status    uint8
}

// TransactionmodelAuction is an auto generated low-level Go binding around an user-defined struct.
type TransactionmodelAuction struct {
	StartedAt *big.Int
	EndedAt   *big.Int
	TokenIds  []*big.Int
}

// AuctionContractMetaData contains all meta data concerning the AuctionContract contract.
var AuctionContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EndAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RegisterNftForAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ScheduleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SellNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StartAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060ba80601a5f395ff3fe6080604052348015600e575f5ffd5b5060043610604e575f3560e01c80630fb28c6e14605257806356ef43f114605a5780636a7798ee14606257806394efd7c614606a578063d295ebf0146072575b5f5ffd5b6058607a565b005b6060607c565b005b6068607e565b005b60706080565b005b60786082565b005b565b565b565b565b56fea26469706673582212209e4f82dc25922513cd7eadcef4059e3375b671dbb5015eb0b0e0dd798ae49f7b64736f6c634300081d0033",
}

// AuctionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionContractMetaData.ABI instead.
var AuctionContractABI = AuctionContractMetaData.ABI

// AuctionContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionContractMetaData.Bin instead.
var AuctionContractBin = AuctionContractMetaData.Bin

// DeployAuctionContract deploys a new Ethereum contract, binding an instance of AuctionContract to it.
func DeployAuctionContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuctionContract, error) {
	parsed, err := AuctionContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionContract{AuctionContractCaller: AuctionContractCaller{contract: contract}, AuctionContractTransactor: AuctionContractTransactor{contract: contract}, AuctionContractFilterer: AuctionContractFilterer{contract: contract}}, nil
}

// AuctionContract is an auto generated Go binding around an Ethereum contract.
type AuctionContract struct {
	AuctionContractCaller     // Read-only binding to the contract
	AuctionContractTransactor // Write-only binding to the contract
	AuctionContractFilterer   // Log filterer for contract events
}

// AuctionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionContractSession struct {
	Contract     *AuctionContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionContractCallerSession struct {
	Contract *AuctionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AuctionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionContractTransactorSession struct {
	Contract     *AuctionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AuctionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionContractRaw struct {
	Contract *AuctionContract // Generic contract binding to access the raw methods on
}

// AuctionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionContractCallerRaw struct {
	Contract *AuctionContractCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionContractTransactorRaw struct {
	Contract *AuctionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionContract creates a new instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContract(address common.Address, backend bind.ContractBackend) (*AuctionContract, error) {
	contract, err := bindAuctionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionContract{AuctionContractCaller: AuctionContractCaller{contract: contract}, AuctionContractTransactor: AuctionContractTransactor{contract: contract}, AuctionContractFilterer: AuctionContractFilterer{contract: contract}}, nil
}

// NewAuctionContractCaller creates a new read-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractCaller(address common.Address, caller bind.ContractCaller) (*AuctionContractCaller, error) {
	contract, err := bindAuctionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractCaller{contract: contract}, nil
}

// NewAuctionContractTransactor creates a new write-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionContractTransactor, error) {
	contract, err := bindAuctionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractTransactor{contract: contract}, nil
}

// NewAuctionContractFilterer creates a new log filterer instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionContractFilterer, error) {
	contract, err := bindAuctionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionContractFilterer{contract: contract}, nil
}

// bindAuctionContract binds a generic wrapper to an already deployed contract.
func bindAuctionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.AuctionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transact(opts, method, params...)
}

// EndAuction is a paid mutator transaction binding the contract method 0x6a7798ee.
//
// Solidity: function EndAuction() returns()
func (_AuctionContract *AuctionContractTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "EndAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0x6a7798ee.
//
// Solidity: function EndAuction() returns()
func (_AuctionContract *AuctionContractSession) EndAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.EndAuction(&_AuctionContract.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0x6a7798ee.
//
// Solidity: function EndAuction() returns()
func (_AuctionContract *AuctionContractTransactorSession) EndAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.EndAuction(&_AuctionContract.TransactOpts)
}

// RegisterNftForAuction is a paid mutator transaction binding the contract method 0x94efd7c6.
//
// Solidity: function RegisterNftForAuction() returns()
func (_AuctionContract *AuctionContractTransactor) RegisterNftForAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "RegisterNftForAuction")
}

// RegisterNftForAuction is a paid mutator transaction binding the contract method 0x94efd7c6.
//
// Solidity: function RegisterNftForAuction() returns()
func (_AuctionContract *AuctionContractSession) RegisterNftForAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.RegisterNftForAuction(&_AuctionContract.TransactOpts)
}

// RegisterNftForAuction is a paid mutator transaction binding the contract method 0x94efd7c6.
//
// Solidity: function RegisterNftForAuction() returns()
func (_AuctionContract *AuctionContractTransactorSession) RegisterNftForAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.RegisterNftForAuction(&_AuctionContract.TransactOpts)
}

// ScheduleAuction is a paid mutator transaction binding the contract method 0x0fb28c6e.
//
// Solidity: function ScheduleAuction() returns()
func (_AuctionContract *AuctionContractTransactor) ScheduleAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "ScheduleAuction")
}

// ScheduleAuction is a paid mutator transaction binding the contract method 0x0fb28c6e.
//
// Solidity: function ScheduleAuction() returns()
func (_AuctionContract *AuctionContractSession) ScheduleAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.ScheduleAuction(&_AuctionContract.TransactOpts)
}

// ScheduleAuction is a paid mutator transaction binding the contract method 0x0fb28c6e.
//
// Solidity: function ScheduleAuction() returns()
func (_AuctionContract *AuctionContractTransactorSession) ScheduleAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.ScheduleAuction(&_AuctionContract.TransactOpts)
}

// SellNft is a paid mutator transaction binding the contract method 0x56ef43f1.
//
// Solidity: function SellNft() returns()
func (_AuctionContract *AuctionContractTransactor) SellNft(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "SellNft")
}

// SellNft is a paid mutator transaction binding the contract method 0x56ef43f1.
//
// Solidity: function SellNft() returns()
func (_AuctionContract *AuctionContractSession) SellNft() (*types.Transaction, error) {
	return _AuctionContract.Contract.SellNft(&_AuctionContract.TransactOpts)
}

// SellNft is a paid mutator transaction binding the contract method 0x56ef43f1.
//
// Solidity: function SellNft() returns()
func (_AuctionContract *AuctionContractTransactorSession) SellNft() (*types.Transaction, error) {
	return _AuctionContract.Contract.SellNft(&_AuctionContract.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0xd295ebf0.
//
// Solidity: function StartAuction() returns()
func (_AuctionContract *AuctionContractTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "StartAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0xd295ebf0.
//
// Solidity: function StartAuction() returns()
func (_AuctionContract *AuctionContractSession) StartAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.StartAuction(&_AuctionContract.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0xd295ebf0.
//
// Solidity: function StartAuction() returns()
func (_AuctionContract *AuctionContractTransactorSession) StartAuction() (*types.Transaction, error) {
	return _AuctionContract.Contract.StartAuction(&_AuctionContract.TransactOpts)
}

// NftContractMetaData contains all meta data concerning the NftContract contract.
var NftContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506108588061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80630788370314610038578063973e7e4014610054575b5f5ffd5b610052600480360381019061004d9190610350565b610070565b005b61006e60048036038101906100699190610350565b610280565b005b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b370317836040518263ffffffff1660e01b81526004016100cb919061038a565b5f60405180830381865afa1580156100e5573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061010d9190610609565b90505f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663aef15d376040518163ffffffff1660e01b81526004016020604051808303815f875af115801561017a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061019e9190610650565b90505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663512b274d826040518060800160405280865f0151815260200186602001518152602001866040015173ffffffffffffffffffffffffffffffffffffffff1681526020015f600481111561022e5761022d61067b565b5b8152506040518363ffffffff1660e01b815260040161024e9291906107be565b5f604051808303815f87803b158015610265575f5ffd5b505af1158015610277573d5f5f3e3d5ffd5b50505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663abb1a5038260016040518363ffffffff1660e01b81526004016102dc9291906107fb565b5f604051808303815f87803b1580156102f3575f5ffd5b505af1158015610305573d5f5f3e3d5ffd5b5050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61032f8161031d565b8114610339575f5ffd5b50565b5f8135905061034a81610326565b92915050565b5f6020828403121561036557610364610315565b5b5f6103728482850161033c565b91505092915050565b6103848161031d565b82525050565b5f60208201905061039d5f83018461037b565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103ed826103a7565b810181811067ffffffffffffffff8211171561040c5761040b6103b7565b5b80604052505050565b5f61041e61030c565b905061042a82826103e4565b919050565b5f5ffd5b5f8151905061044181610326565b92915050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff821115610469576104686103b7565b5b610472826103a7565b9050602081019050919050565b8281835e5f83830152505050565b5f61049f61049a8461044f565b610415565b9050828152602081018484840111156104bb576104ba61044b565b5b6104c684828561047f565b509392505050565b5f82601f8301126104e2576104e1610447565b5b81516104f284826020860161048d565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610524826104fb565b9050919050565b6105348161051a565b811461053e575f5ffd5b50565b5f8151905061054f8161052b565b92915050565b60058110610561575f5ffd5b50565b5f8151905061057281610555565b92915050565b5f6080828403121561058d5761058c6103a3565b5b6105976080610415565b90505f6105a684828501610433565b5f83015250602082015167ffffffffffffffff8111156105c9576105c861042f565b5b6105d5848285016104ce565b60208301525060406105e984828501610541565b60408301525060606105fd84828501610564565b60608301525092915050565b5f6020828403121561061e5761061d610315565b5b5f82015167ffffffffffffffff81111561063b5761063a610319565b5b61064784828501610578565b91505092915050565b5f6020828403121561066557610664610315565b5b5f61067284828501610433565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6106b18161031d565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f6106db826106b7565b6106e581856106c1565b93506106f581856020860161047f565b6106fe816103a7565b840191505092915050565b6107128161051a565b82525050565b600581106107295761072861067b565b5b50565b5f81905061073982610718565b919050565b5f6107488261072c565b9050919050565b6107588161073e565b82525050565b5f608083015f8301516107735f8601826106a8565b506020830151848203602086015261078b82826106d1565b91505060408301516107a06040860182610709565b5060608301516107b3606086018261074f565b508091505092915050565b5f6040820190506107d15f83018561037b565b81810360208301526107e3818461075e565b90509392505050565b6107f58161073e565b82525050565b5f60408201905061080e5f83018561037b565b61081b60208301846107ec565b939250505056fea2646970667358221220465c2504a4980c9bcabc63149f76db5e501a9a6d60be23bcf60b1b93a9f1aa4864736f6c634300081d0033",
}

// NftContractABI is the input ABI used to generate the binding from.
// Deprecated: Use NftContractMetaData.ABI instead.
var NftContractABI = NftContractMetaData.ABI

// NftContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NftContractMetaData.Bin instead.
var NftContractBin = NftContractMetaData.Bin

// DeployNftContract deploys a new Ethereum contract, binding an instance of NftContract to it.
func DeployNftContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NftContract, error) {
	parsed, err := NftContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NftContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NftContract{NftContractCaller: NftContractCaller{contract: contract}, NftContractTransactor: NftContractTransactor{contract: contract}, NftContractFilterer: NftContractFilterer{contract: contract}}, nil
}

// NftContract is an auto generated Go binding around an Ethereum contract.
type NftContract struct {
	NftContractCaller     // Read-only binding to the contract
	NftContractTransactor // Write-only binding to the contract
	NftContractFilterer   // Log filterer for contract events
}

// NftContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftContractSession struct {
	Contract     *NftContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftContractCallerSession struct {
	Contract *NftContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NftContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftContractTransactorSession struct {
	Contract     *NftContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NftContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftContractRaw struct {
	Contract *NftContract // Generic contract binding to access the raw methods on
}

// NftContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftContractCallerRaw struct {
	Contract *NftContractCaller // Generic read-only contract binding to access the raw methods on
}

// NftContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftContractTransactorRaw struct {
	Contract *NftContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftContract creates a new instance of NftContract, bound to a specific deployed contract.
func NewNftContract(address common.Address, backend bind.ContractBackend) (*NftContract, error) {
	contract, err := bindNftContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftContract{NftContractCaller: NftContractCaller{contract: contract}, NftContractTransactor: NftContractTransactor{contract: contract}, NftContractFilterer: NftContractFilterer{contract: contract}}, nil
}

// NewNftContractCaller creates a new read-only instance of NftContract, bound to a specific deployed contract.
func NewNftContractCaller(address common.Address, caller bind.ContractCaller) (*NftContractCaller, error) {
	contract, err := bindNftContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftContractCaller{contract: contract}, nil
}

// NewNftContractTransactor creates a new write-only instance of NftContract, bound to a specific deployed contract.
func NewNftContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NftContractTransactor, error) {
	contract, err := bindNftContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftContractTransactor{contract: contract}, nil
}

// NewNftContractFilterer creates a new log filterer instance of NftContract, bound to a specific deployed contract.
func NewNftContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NftContractFilterer, error) {
	contract, err := bindNftContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftContractFilterer{contract: contract}, nil
}

// bindNftContract binds a generic wrapper to an already deployed contract.
func bindNftContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftContract *NftContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftContract.Contract.NftContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftContract *NftContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftContract.Contract.NftContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftContract *NftContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftContract.Contract.NftContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftContract *NftContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftContract *NftContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftContract *NftContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftContract.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0x973e7e40.
//
// Solidity: function Activate(uint256 tokenId) returns()
func (_NftContract *NftContractTransactor) Activate(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NftContract.contract.Transact(opts, "Activate", tokenId)
}

// Activate is a paid mutator transaction binding the contract method 0x973e7e40.
//
// Solidity: function Activate(uint256 tokenId) returns()
func (_NftContract *NftContractSession) Activate(tokenId *big.Int) (*types.Transaction, error) {
	return _NftContract.Contract.Activate(&_NftContract.TransactOpts, tokenId)
}

// Activate is a paid mutator transaction binding the contract method 0x973e7e40.
//
// Solidity: function Activate(uint256 tokenId) returns()
func (_NftContract *NftContractTransactorSession) Activate(tokenId *big.Int) (*types.Transaction, error) {
	return _NftContract.Contract.Activate(&_NftContract.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x07883703.
//
// Solidity: function Mint(uint256 requestId) returns()
func (_NftContract *NftContractTransactor) Mint(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _NftContract.contract.Transact(opts, "Mint", requestId)
}

// Mint is a paid mutator transaction binding the contract method 0x07883703.
//
// Solidity: function Mint(uint256 requestId) returns()
func (_NftContract *NftContractSession) Mint(requestId *big.Int) (*types.Transaction, error) {
	return _NftContract.Contract.Mint(&_NftContract.TransactOpts, requestId)
}

// Mint is a paid mutator transaction binding the contract method 0x07883703.
//
// Solidity: function Mint(uint256 requestId) returns()
func (_NftContract *NftContractTransactorSession) Mint(requestId *big.Int) (*types.Transaction, error) {
	return _NftContract.Contract.Mint(&_NftContract.TransactOpts, requestId)
}

// OnBoardingContractMetaData contains all meta data concerning the OnBoardingContract contract.
var OnBoardingContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"OnBoardEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumorgmodel.Role\",\"name\":\"_role\",\"type\":\"uint8\"}],\"name\":\"OnBoardOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102ee8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063d339fd571461002d575b5f5ffd5b61004760048036038101906100429190610190565b610049565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631b56f91383836040518363ffffffff1660e01b81526004016100a4929190610250565b5f604051808303815f87803b1580156100bb575f5ffd5b505af11580156100cd573d5f5f3e3d5ffd5b505050507f97fcd11107fe7ad17c3eec405f9097a62e72ea036cdd4c07cdf9b53f8def0d64826001604051610103929190610291565b60405180910390a15050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61013c82610113565b9050919050565b61014c81610132565b8114610156575f5ffd5b50565b5f8135905061016781610143565b92915050565b60048110610179575f5ffd5b50565b5f8135905061018a8161016d565b92915050565b5f5f604083850312156101a6576101a561010f565b5b5f6101b385828601610159565b92505060206101c48582860161017c565b9150509250929050565b6101d781610132565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6004811061021b5761021a6101dd565b5b50565b5f81905061022b8261020a565b919050565b5f61023a8261021e565b9050919050565b61024a81610230565b82525050565b5f6040820190506102635f8301856101ce565b6102706020830184610241565b9392505050565b5f8115159050919050565b61028b81610277565b82525050565b5f6040820190506102a45f8301856101ce565b6102b16020830184610282565b939250505056fea264697066735822122090085c359fbf52a794f7e99ede544767bf1b6d6dcad189aab0e517a25c3564bf64736f6c634300081d0033",
}

// OnBoardingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OnBoardingContractMetaData.ABI instead.
var OnBoardingContractABI = OnBoardingContractMetaData.ABI

// OnBoardingContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnBoardingContractMetaData.Bin instead.
var OnBoardingContractBin = OnBoardingContractMetaData.Bin

// DeployOnBoardingContract deploys a new Ethereum contract, binding an instance of OnBoardingContract to it.
func DeployOnBoardingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnBoardingContract, error) {
	parsed, err := OnBoardingContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnBoardingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnBoardingContract{OnBoardingContractCaller: OnBoardingContractCaller{contract: contract}, OnBoardingContractTransactor: OnBoardingContractTransactor{contract: contract}, OnBoardingContractFilterer: OnBoardingContractFilterer{contract: contract}}, nil
}

// OnBoardingContract is an auto generated Go binding around an Ethereum contract.
type OnBoardingContract struct {
	OnBoardingContractCaller     // Read-only binding to the contract
	OnBoardingContractTransactor // Write-only binding to the contract
	OnBoardingContractFilterer   // Log filterer for contract events
}

// OnBoardingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnBoardingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnBoardingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnBoardingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnBoardingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnBoardingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnBoardingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnBoardingContractSession struct {
	Contract     *OnBoardingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OnBoardingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnBoardingContractCallerSession struct {
	Contract *OnBoardingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OnBoardingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnBoardingContractTransactorSession struct {
	Contract     *OnBoardingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OnBoardingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnBoardingContractRaw struct {
	Contract *OnBoardingContract // Generic contract binding to access the raw methods on
}

// OnBoardingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnBoardingContractCallerRaw struct {
	Contract *OnBoardingContractCaller // Generic read-only contract binding to access the raw methods on
}

// OnBoardingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnBoardingContractTransactorRaw struct {
	Contract *OnBoardingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnBoardingContract creates a new instance of OnBoardingContract, bound to a specific deployed contract.
func NewOnBoardingContract(address common.Address, backend bind.ContractBackend) (*OnBoardingContract, error) {
	contract, err := bindOnBoardingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnBoardingContract{OnBoardingContractCaller: OnBoardingContractCaller{contract: contract}, OnBoardingContractTransactor: OnBoardingContractTransactor{contract: contract}, OnBoardingContractFilterer: OnBoardingContractFilterer{contract: contract}}, nil
}

// NewOnBoardingContractCaller creates a new read-only instance of OnBoardingContract, bound to a specific deployed contract.
func NewOnBoardingContractCaller(address common.Address, caller bind.ContractCaller) (*OnBoardingContractCaller, error) {
	contract, err := bindOnBoardingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnBoardingContractCaller{contract: contract}, nil
}

// NewOnBoardingContractTransactor creates a new write-only instance of OnBoardingContract, bound to a specific deployed contract.
func NewOnBoardingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OnBoardingContractTransactor, error) {
	contract, err := bindOnBoardingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnBoardingContractTransactor{contract: contract}, nil
}

// NewOnBoardingContractFilterer creates a new log filterer instance of OnBoardingContract, bound to a specific deployed contract.
func NewOnBoardingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OnBoardingContractFilterer, error) {
	contract, err := bindOnBoardingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnBoardingContractFilterer{contract: contract}, nil
}

// bindOnBoardingContract binds a generic wrapper to an already deployed contract.
func bindOnBoardingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnBoardingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnBoardingContract *OnBoardingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnBoardingContract.Contract.OnBoardingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnBoardingContract *OnBoardingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.OnBoardingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnBoardingContract *OnBoardingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.OnBoardingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnBoardingContract *OnBoardingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnBoardingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnBoardingContract *OnBoardingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnBoardingContract *OnBoardingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.contract.Transact(opts, method, params...)
}

// OnBoardOrg is a paid mutator transaction binding the contract method 0xd339fd57.
//
// Solidity: function OnBoardOrg(address _address, uint8 _role) returns()
func (_OnBoardingContract *OnBoardingContractTransactor) OnBoardOrg(opts *bind.TransactOpts, _address common.Address, _role uint8) (*types.Transaction, error) {
	return _OnBoardingContract.contract.Transact(opts, "OnBoardOrg", _address, _role)
}

// OnBoardOrg is a paid mutator transaction binding the contract method 0xd339fd57.
//
// Solidity: function OnBoardOrg(address _address, uint8 _role) returns()
func (_OnBoardingContract *OnBoardingContractSession) OnBoardOrg(_address common.Address, _role uint8) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.OnBoardOrg(&_OnBoardingContract.TransactOpts, _address, _role)
}

// OnBoardOrg is a paid mutator transaction binding the contract method 0xd339fd57.
//
// Solidity: function OnBoardOrg(address _address, uint8 _role) returns()
func (_OnBoardingContract *OnBoardingContractTransactorSession) OnBoardOrg(_address common.Address, _role uint8) (*types.Transaction, error) {
	return _OnBoardingContract.Contract.OnBoardOrg(&_OnBoardingContract.TransactOpts, _address, _role)
}

// OnBoardingContractOnBoardEventIterator is returned from FilterOnBoardEvent and is used to iterate over the raw logs and unpacked data for OnBoardEvent events raised by the OnBoardingContract contract.
type OnBoardingContractOnBoardEventIterator struct {
	Event *OnBoardingContractOnBoardEvent // Event containing the contract specifics and raw log

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
func (it *OnBoardingContractOnBoardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnBoardingContractOnBoardEvent)
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
		it.Event = new(OnBoardingContractOnBoardEvent)
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
func (it *OnBoardingContractOnBoardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnBoardingContractOnBoardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnBoardingContractOnBoardEvent represents a OnBoardEvent event raised by the OnBoardingContract contract.
type OnBoardingContractOnBoardEvent struct {
	Address common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnBoardEvent is a free log retrieval operation binding the contract event 0x97fcd11107fe7ad17c3eec405f9097a62e72ea036cdd4c07cdf9b53f8def0d64.
//
// Solidity: event OnBoardEvent(address _address, bool status)
func (_OnBoardingContract *OnBoardingContractFilterer) FilterOnBoardEvent(opts *bind.FilterOpts) (*OnBoardingContractOnBoardEventIterator, error) {

	logs, sub, err := _OnBoardingContract.contract.FilterLogs(opts, "OnBoardEvent")
	if err != nil {
		return nil, err
	}
	return &OnBoardingContractOnBoardEventIterator{contract: _OnBoardingContract.contract, event: "OnBoardEvent", logs: logs, sub: sub}, nil
}

// WatchOnBoardEvent is a free log subscription operation binding the contract event 0x97fcd11107fe7ad17c3eec405f9097a62e72ea036cdd4c07cdf9b53f8def0d64.
//
// Solidity: event OnBoardEvent(address _address, bool status)
func (_OnBoardingContract *OnBoardingContractFilterer) WatchOnBoardEvent(opts *bind.WatchOpts, sink chan<- *OnBoardingContractOnBoardEvent) (event.Subscription, error) {

	logs, sub, err := _OnBoardingContract.contract.WatchLogs(opts, "OnBoardEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnBoardingContractOnBoardEvent)
				if err := _OnBoardingContract.contract.UnpackLog(event, "OnBoardEvent", log); err != nil {
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

// ParseOnBoardEvent is a log parse operation binding the contract event 0x97fcd11107fe7ad17c3eec405f9097a62e72ea036cdd4c07cdf9b53f8def0d64.
//
// Solidity: event OnBoardEvent(address _address, bool status)
func (_OnBoardingContract *OnBoardingContractFilterer) ParseOnBoardEvent(log types.Log) (*OnBoardingContractOnBoardEvent, error) {
	event := new(OnBoardingContractOnBoardEvent)
	if err := _OnBoardingContract.contract.UnpackLog(event, "OnBoardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrgStoreMetaData contains all meta data concerning the OrgStore contract.
var OrgStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetOrg\",\"outputs\":[{\"internalType\":\"enumorgmodel.Role\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumorgmodel.Role\",\"name\":\"_role\",\"type\":\"uint8\"}],\"name\":\"SaveOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102e88061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80631b56f91314610038578063ce3715a914610054575b5f5ffd5b610052600480360381019061004d91906101bd565b610084565b005b61006e600480360381019061006991906101fb565b6100eb565b60405161007b9190610299565b60405180910390f35b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908360038111156100e2576100e1610226565b5b02179055505050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61016982610140565b9050919050565b6101798161015f565b8114610183575f5ffd5b50565b5f8135905061019481610170565b92915050565b600481106101a6575f5ffd5b50565b5f813590506101b78161019a565b92915050565b5f5f604083850312156101d3576101d261013c565b5b5f6101e085828601610186565b92505060206101f1858286016101a9565b9150509250929050565b5f602082840312156102105761020f61013c565b5b5f61021d84828501610186565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6004811061026457610263610226565b5b50565b5f81905061027482610253565b919050565b5f61028382610267565b9050919050565b61029381610279565b82525050565b5f6020820190506102ac5f83018461028a565b9291505056fea264697066735822122007cb7dda37238b052ebe2ec9d1adca84aa601145cc227f7478c535949f2415ba64736f6c634300081d0033",
}

// OrgStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use OrgStoreMetaData.ABI instead.
var OrgStoreABI = OrgStoreMetaData.ABI

// OrgStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrgStoreMetaData.Bin instead.
var OrgStoreBin = OrgStoreMetaData.Bin

// DeployOrgStore deploys a new Ethereum contract, binding an instance of OrgStore to it.
func DeployOrgStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrgStore, error) {
	parsed, err := OrgStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrgStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrgStore{OrgStoreCaller: OrgStoreCaller{contract: contract}, OrgStoreTransactor: OrgStoreTransactor{contract: contract}, OrgStoreFilterer: OrgStoreFilterer{contract: contract}}, nil
}

// OrgStore is an auto generated Go binding around an Ethereum contract.
type OrgStore struct {
	OrgStoreCaller     // Read-only binding to the contract
	OrgStoreTransactor // Write-only binding to the contract
	OrgStoreFilterer   // Log filterer for contract events
}

// OrgStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgStoreSession struct {
	Contract     *OrgStore         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgStoreCallerSession struct {
	Contract *OrgStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OrgStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgStoreTransactorSession struct {
	Contract     *OrgStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OrgStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgStoreRaw struct {
	Contract *OrgStore // Generic contract binding to access the raw methods on
}

// OrgStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgStoreCallerRaw struct {
	Contract *OrgStoreCaller // Generic read-only contract binding to access the raw methods on
}

// OrgStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgStoreTransactorRaw struct {
	Contract *OrgStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgStore creates a new instance of OrgStore, bound to a specific deployed contract.
func NewOrgStore(address common.Address, backend bind.ContractBackend) (*OrgStore, error) {
	contract, err := bindOrgStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrgStore{OrgStoreCaller: OrgStoreCaller{contract: contract}, OrgStoreTransactor: OrgStoreTransactor{contract: contract}, OrgStoreFilterer: OrgStoreFilterer{contract: contract}}, nil
}

// NewOrgStoreCaller creates a new read-only instance of OrgStore, bound to a specific deployed contract.
func NewOrgStoreCaller(address common.Address, caller bind.ContractCaller) (*OrgStoreCaller, error) {
	contract, err := bindOrgStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgStoreCaller{contract: contract}, nil
}

// NewOrgStoreTransactor creates a new write-only instance of OrgStore, bound to a specific deployed contract.
func NewOrgStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgStoreTransactor, error) {
	contract, err := bindOrgStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgStoreTransactor{contract: contract}, nil
}

// NewOrgStoreFilterer creates a new log filterer instance of OrgStore, bound to a specific deployed contract.
func NewOrgStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgStoreFilterer, error) {
	contract, err := bindOrgStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgStoreFilterer{contract: contract}, nil
}

// bindOrgStore binds a generic wrapper to an already deployed contract.
func bindOrgStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrgStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgStore *OrgStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrgStore.Contract.OrgStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgStore *OrgStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgStore.Contract.OrgStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgStore *OrgStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgStore.Contract.OrgStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgStore *OrgStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrgStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgStore *OrgStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgStore *OrgStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgStore.Contract.contract.Transact(opts, method, params...)
}

// GetOrg is a free data retrieval call binding the contract method 0xce3715a9.
//
// Solidity: function GetOrg(address _address) view returns(uint8)
func (_OrgStore *OrgStoreCaller) GetOrg(opts *bind.CallOpts, _address common.Address) (uint8, error) {
	var out []interface{}
	err := _OrgStore.contract.Call(opts, &out, "GetOrg", _address)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOrg is a free data retrieval call binding the contract method 0xce3715a9.
//
// Solidity: function GetOrg(address _address) view returns(uint8)
func (_OrgStore *OrgStoreSession) GetOrg(_address common.Address) (uint8, error) {
	return _OrgStore.Contract.GetOrg(&_OrgStore.CallOpts, _address)
}

// GetOrg is a free data retrieval call binding the contract method 0xce3715a9.
//
// Solidity: function GetOrg(address _address) view returns(uint8)
func (_OrgStore *OrgStoreCallerSession) GetOrg(_address common.Address) (uint8, error) {
	return _OrgStore.Contract.GetOrg(&_OrgStore.CallOpts, _address)
}

// SaveOrg is a paid mutator transaction binding the contract method 0x1b56f913.
//
// Solidity: function SaveOrg(address _address, uint8 _role) returns()
func (_OrgStore *OrgStoreTransactor) SaveOrg(opts *bind.TransactOpts, _address common.Address, _role uint8) (*types.Transaction, error) {
	return _OrgStore.contract.Transact(opts, "SaveOrg", _address, _role)
}

// SaveOrg is a paid mutator transaction binding the contract method 0x1b56f913.
//
// Solidity: function SaveOrg(address _address, uint8 _role) returns()
func (_OrgStore *OrgStoreSession) SaveOrg(_address common.Address, _role uint8) (*types.Transaction, error) {
	return _OrgStore.Contract.SaveOrg(&_OrgStore.TransactOpts, _address, _role)
}

// SaveOrg is a paid mutator transaction binding the contract method 0x1b56f913.
//
// Solidity: function SaveOrg(address _address, uint8 _role) returns()
func (_OrgStore *OrgStoreTransactorSession) SaveOrg(_address common.Address, _role uint8) (*types.Transaction, error) {
	return _OrgStore.Contract.SaveOrg(&_OrgStore.TransactOpts, _address, _role)
}

// RequestContractMetaData contains all meta data concerning the RequestContract contract.
var RequestContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ApprovedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SavedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"queryRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumrequestmodel.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structrequestmodel.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumrequestmodel.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structrequestmodel.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"saveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610bb28061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80638fb0e30b14610043578063934a68a21461005f578063d7d1bbdb1461008f575b5f5ffd5b61005d60048036038101906100589190610703565b6100ab565b005b6100796004803603810190610074919061075d565b6101a0565b60405161008691906108d9565b60405180910390f35b6100a960048036038101906100a4919061075d565b61024a565b005b5f816060019060048111156100c3576100c2610806565b5b908160048111156100d7576100d6610806565b5b815250505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c997fe8b83836040518363ffffffff1660e01b8152600401610136929190610908565b5f604051808303815f87803b15801561014d575f5ffd5b505af115801561015f573d5f5f3e3d5ffd5b50505050817f7a6c713447bfbd5445c5eb63352ce75b252613712130c0056d11db3d76d9467b60016040516101949190610950565b60405180910390a25050565b6101a8610420565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b370317836040518263ffffffff1660e01b81526004016102019190610969565b5f60405180830381865afa15801561021b573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102439190610abd565b9050919050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b370317836040518263ffffffff1660e01b81526004016102a49190610969565b5f60405180830381865afa1580156102be573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102e69190610abd565b90505f820361032a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032190610b5e565b60405180910390fd5b60018160600190600481111561034357610342610806565b5b9081600481111561035757610356610806565b5b815250505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c997fe8b83836040518363ffffffff1660e01b81526004016103b6929190610908565b5f604051808303815f87803b1580156103cd575f5ffd5b505af11580156103df573d5f5f3e3d5ffd5b50505050817fd77f0e83ebaad1ffb1d409decd47fb1fb5bdd329b6f346d341b0b4533edb9d9060016040516104149190610950565b60405180910390a25050565b60405180608001604052805f8152602001606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f600481111561046757610466610806565b5b81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6104908161047e565b811461049a575f5ffd5b50565b5f813590506104ab81610487565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104fb826104b5565b810181811067ffffffffffffffff8211171561051a576105196104c5565b5b80604052505050565b5f61052c61046d565b905061053882826104f2565b919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f67ffffffffffffffff821115610563576105626104c5565b5b61056c826104b5565b9050602081019050919050565b828183375f83830152505050565b5f61059961059484610549565b610523565b9050828152602081018484840111156105b5576105b4610545565b5b6105c0848285610579565b509392505050565b5f82601f8301126105dc576105db610541565b5b81356105ec848260208601610587565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61061e826105f5565b9050919050565b61062e81610614565b8114610638575f5ffd5b50565b5f8135905061064981610625565b92915050565b6005811061065b575f5ffd5b50565b5f8135905061066c8161064f565b92915050565b5f60808284031215610687576106866104b1565b5b6106916080610523565b90505f6106a08482850161049d565b5f83015250602082013567ffffffffffffffff8111156106c3576106c261053d565b5b6106cf848285016105c8565b60208301525060406106e38482850161063b565b60408301525060606106f78482850161065e565b60608301525092915050565b5f5f6040838503121561071957610718610476565b5b5f6107268582860161049d565b925050602083013567ffffffffffffffff8111156107475761074661047a565b5b61075385828601610672565b9150509250929050565b5f6020828403121561077257610771610476565b5b5f61077f8482850161049d565b91505092915050565b6107918161047e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6107c982610797565b6107d381856107a1565b93506107e38185602086016107b1565b6107ec816104b5565b840191505092915050565b61080081610614565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061084457610843610806565b5b50565b5f81905061085482610833565b919050565b5f61086382610847565b9050919050565b61087381610859565b82525050565b5f608083015f83015161088e5f860182610788565b50602083015184820360208601526108a682826107bf565b91505060408301516108bb60408601826107f7565b5060608301516108ce606086018261086a565b508091505092915050565b5f6020820190508181035f8301526108f18184610879565b905092915050565b6109028161047e565b82525050565b5f60408201905061091b5f8301856108f9565b818103602083015261092d8184610879565b90509392505050565b5f8115159050919050565b61094a81610936565b82525050565b5f6020820190506109635f830184610941565b92915050565b5f60208201905061097c5f8301846108f9565b92915050565b5f8151905061099081610487565b92915050565b5f6109a86109a384610549565b610523565b9050828152602081018484840111156109c4576109c3610545565b5b6109cf8482856107b1565b509392505050565b5f82601f8301126109eb576109ea610541565b5b81516109fb848260208601610996565b91505092915050565b5f81519050610a1281610625565b92915050565b5f81519050610a268161064f565b92915050565b5f60808284031215610a4157610a406104b1565b5b610a4b6080610523565b90505f610a5a84828501610982565b5f83015250602082015167ffffffffffffffff811115610a7d57610a7c61053d565b5b610a89848285016109d7565b6020830152506040610a9d84828501610a04565b6040830152506060610ab184828501610a18565b60608301525092915050565b5f60208284031215610ad257610ad1610476565b5b5f82015167ffffffffffffffff811115610aef57610aee61047a565b5b610afb84828501610a2c565b91505092915050565b5f82825260208201905092915050565b7f52657175657374206e6f7420666f756e640000000000000000000000000000005f82015250565b5f610b48601183610b04565b9150610b5382610b14565b602082019050919050565b5f6020820190508181035f830152610b7581610b3c565b905091905056fea26469706673582212209218af9df1077ce0de4502704c258c9d11edcea14c3f64518135f503914187a064736f6c634300081d0033",
}

// RequestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestContractMetaData.ABI instead.
var RequestContractABI = RequestContractMetaData.ABI

// RequestContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestContractMetaData.Bin instead.
var RequestContractBin = RequestContractMetaData.Bin

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := RequestContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}, RequestContractFilterer: RequestContractFilterer{contract: contract}}, nil
}

// RequestContract is an auto generated Go binding around an Ethereum contract.
type RequestContract struct {
	RequestContractCaller     // Read-only binding to the contract
	RequestContractTransactor // Write-only binding to the contract
	RequestContractFilterer   // Log filterer for contract events
}

// RequestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestContractSession struct {
	Contract     *RequestContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestContractCallerSession struct {
	Contract *RequestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RequestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestContractTransactorSession struct {
	Contract     *RequestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RequestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestContractRaw struct {
	Contract *RequestContract // Generic contract binding to access the raw methods on
}

// RequestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestContractCallerRaw struct {
	Contract *RequestContractCaller // Generic read-only contract binding to access the raw methods on
}

// RequestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestContractTransactorRaw struct {
	Contract *RequestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestContract creates a new instance of RequestContract, bound to a specific deployed contract.
func NewRequestContract(address common.Address, backend bind.ContractBackend) (*RequestContract, error) {
	contract, err := bindRequestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}, RequestContractFilterer: RequestContractFilterer{contract: contract}}, nil
}

// NewRequestContractCaller creates a new read-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractCaller(address common.Address, caller bind.ContractCaller) (*RequestContractCaller, error) {
	contract, err := bindRequestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestContractCaller{contract: contract}, nil
}

// NewRequestContractTransactor creates a new write-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestContractTransactor, error) {
	contract, err := bindRequestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestContractTransactor{contract: contract}, nil
}

// NewRequestContractFilterer creates a new log filterer instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestContractFilterer, error) {
	contract, err := bindRequestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestContractFilterer{contract: contract}, nil
}

// bindRequestContract binds a generic wrapper to an already deployed contract.
func bindRequestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.RequestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transact(opts, method, params...)
}

// QueryRequest is a free data retrieval call binding the contract method 0x934a68a2.
//
// Solidity: function queryRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_RequestContract *RequestContractCaller) QueryRequest(opts *bind.CallOpts, id *big.Int) (RequestmodelRequest, error) {
	var out []interface{}
	err := _RequestContract.contract.Call(opts, &out, "queryRequest", id)

	if err != nil {
		return *new(RequestmodelRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(RequestmodelRequest)).(*RequestmodelRequest)

	return out0, err

}

// QueryRequest is a free data retrieval call binding the contract method 0x934a68a2.
//
// Solidity: function queryRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_RequestContract *RequestContractSession) QueryRequest(id *big.Int) (RequestmodelRequest, error) {
	return _RequestContract.Contract.QueryRequest(&_RequestContract.CallOpts, id)
}

// QueryRequest is a free data retrieval call binding the contract method 0x934a68a2.
//
// Solidity: function queryRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_RequestContract *RequestContractCallerSession) QueryRequest(id *big.Int) (RequestmodelRequest, error) {
	return _RequestContract.Contract.QueryRequest(&_RequestContract.CallOpts, id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 id) returns()
func (_RequestContract *RequestContractTransactor) ApproveRequest(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "approveRequest", id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 id) returns()
func (_RequestContract *RequestContractSession) ApproveRequest(id *big.Int) (*types.Transaction, error) {
	return _RequestContract.Contract.ApproveRequest(&_RequestContract.TransactOpts, id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 id) returns()
func (_RequestContract *RequestContractTransactorSession) ApproveRequest(id *big.Int) (*types.Transaction, error) {
	return _RequestContract.Contract.ApproveRequest(&_RequestContract.TransactOpts, id)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x8fb0e30b.
//
// Solidity: function saveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_RequestContract *RequestContractTransactor) SaveRequest(opts *bind.TransactOpts, id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "saveRequest", id, _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x8fb0e30b.
//
// Solidity: function saveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_RequestContract *RequestContractSession) SaveRequest(id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _RequestContract.Contract.SaveRequest(&_RequestContract.TransactOpts, id, _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x8fb0e30b.
//
// Solidity: function saveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_RequestContract *RequestContractTransactorSession) SaveRequest(id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _RequestContract.Contract.SaveRequest(&_RequestContract.TransactOpts, id, _request)
}

// RequestContractApprovedEventIterator is returned from FilterApprovedEvent and is used to iterate over the raw logs and unpacked data for ApprovedEvent events raised by the RequestContract contract.
type RequestContractApprovedEventIterator struct {
	Event *RequestContractApprovedEvent // Event containing the contract specifics and raw log

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
func (it *RequestContractApprovedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestContractApprovedEvent)
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
		it.Event = new(RequestContractApprovedEvent)
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
func (it *RequestContractApprovedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestContractApprovedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestContractApprovedEvent represents a ApprovedEvent event raised by the RequestContract contract.
type RequestContractApprovedEvent struct {
	Id     *big.Int
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApprovedEvent is a free log retrieval operation binding the contract event 0xd77f0e83ebaad1ffb1d409decd47fb1fb5bdd329b6f346d341b0b4533edb9d90.
//
// Solidity: event ApprovedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) FilterApprovedEvent(opts *bind.FilterOpts, id []*big.Int) (*RequestContractApprovedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RequestContract.contract.FilterLogs(opts, "ApprovedEvent", idRule)
	if err != nil {
		return nil, err
	}
	return &RequestContractApprovedEventIterator{contract: _RequestContract.contract, event: "ApprovedEvent", logs: logs, sub: sub}, nil
}

// WatchApprovedEvent is a free log subscription operation binding the contract event 0xd77f0e83ebaad1ffb1d409decd47fb1fb5bdd329b6f346d341b0b4533edb9d90.
//
// Solidity: event ApprovedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) WatchApprovedEvent(opts *bind.WatchOpts, sink chan<- *RequestContractApprovedEvent, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RequestContract.contract.WatchLogs(opts, "ApprovedEvent", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestContractApprovedEvent)
				if err := _RequestContract.contract.UnpackLog(event, "ApprovedEvent", log); err != nil {
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

// ParseApprovedEvent is a log parse operation binding the contract event 0xd77f0e83ebaad1ffb1d409decd47fb1fb5bdd329b6f346d341b0b4533edb9d90.
//
// Solidity: event ApprovedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) ParseApprovedEvent(log types.Log) (*RequestContractApprovedEvent, error) {
	event := new(RequestContractApprovedEvent)
	if err := _RequestContract.contract.UnpackLog(event, "ApprovedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestContractSavedEventIterator is returned from FilterSavedEvent and is used to iterate over the raw logs and unpacked data for SavedEvent events raised by the RequestContract contract.
type RequestContractSavedEventIterator struct {
	Event *RequestContractSavedEvent // Event containing the contract specifics and raw log

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
func (it *RequestContractSavedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestContractSavedEvent)
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
		it.Event = new(RequestContractSavedEvent)
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
func (it *RequestContractSavedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestContractSavedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestContractSavedEvent represents a SavedEvent event raised by the RequestContract contract.
type RequestContractSavedEvent struct {
	Id     *big.Int
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSavedEvent is a free log retrieval operation binding the contract event 0x7a6c713447bfbd5445c5eb63352ce75b252613712130c0056d11db3d76d9467b.
//
// Solidity: event SavedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) FilterSavedEvent(opts *bind.FilterOpts, id []*big.Int) (*RequestContractSavedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RequestContract.contract.FilterLogs(opts, "SavedEvent", idRule)
	if err != nil {
		return nil, err
	}
	return &RequestContractSavedEventIterator{contract: _RequestContract.contract, event: "SavedEvent", logs: logs, sub: sub}, nil
}

// WatchSavedEvent is a free log subscription operation binding the contract event 0x7a6c713447bfbd5445c5eb63352ce75b252613712130c0056d11db3d76d9467b.
//
// Solidity: event SavedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) WatchSavedEvent(opts *bind.WatchOpts, sink chan<- *RequestContractSavedEvent, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RequestContract.contract.WatchLogs(opts, "SavedEvent", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestContractSavedEvent)
				if err := _RequestContract.contract.UnpackLog(event, "SavedEvent", log); err != nil {
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

// ParseSavedEvent is a log parse operation binding the contract event 0x7a6c713447bfbd5445c5eb63352ce75b252613712130c0056d11db3d76d9467b.
//
// Solidity: event SavedEvent(uint256 indexed id, bool status)
func (_RequestContract *RequestContractFilterer) ParseSavedEvent(log types.Log) (*RequestContractSavedEvent, error) {
	event := new(RequestContractSavedEvent)
	if err := _RequestContract.contract.UnpackLog(event, "SavedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionsStoreMetaData contains all meta data concerning the TransactionsStore contract.
var TransactionsStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structtransactionmodel.Auction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structtransactionmodel.Auction\",\"name\":\"_auction\",\"type\":\"tuple\"}],\"name\":\"SetAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105db8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806330c6343914610038578063f41c5ff714610068575b5f5ffd5b610052600480360381019061004d9190610227565b610084565b60405161005f9190610356565b60405180910390f35b610082600480360381019061007d919061054b565b610117565b005b61008c61015e565b5f5f8381526020019081526020015f206040518060600160405290815f8201548152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561010757602002820191905f5260205f20905b8154815260200190600101908083116100f3575b5050505050815250509050919050565b805f5f8481526020019081526020015f205f820151815f015560208201518160010155604082015181600201908051906020019061015692919061017d565b509050505050565b60405180606001604052805f81526020015f8152602001606081525090565b828054828255905f5260205f209081019282156101b7579160200282015b828111156101b657825182559160200191906001019061019b565b5b5090506101c491906101c8565b5090565b5b808211156101df575f815f9055506001016101c9565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610206816101f4565b8114610210575f5ffd5b50565b5f81359050610221816101fd565b92915050565b5f6020828403121561023c5761023b6101ec565b5b5f61024984828501610213565b91505092915050565b61025b816101f4565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6102958383610252565b60208301905092915050565b5f602082019050919050565b5f6102b782610261565b6102c1818561026b565b93506102cc8361027b565b805f5b838110156102fc5781516102e3888261028a565b97506102ee836102a1565b9250506001810190506102cf565b5085935050505092915050565b5f606083015f83015161031e5f860182610252565b5060208301516103316020860182610252565b506040830151848203604086015261034982826102ad565b9150508091505092915050565b5f6020820190508181035f83015261036e8184610309565b905092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103c08261037a565b810181811067ffffffffffffffff821117156103df576103de61038a565b5b80604052505050565b5f6103f16101e3565b90506103fd82826103b7565b919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff8211156104245761042361038a565b5b602082029050602081019050919050565b5f5ffd5b5f61044b6104468461040a565b6103e8565b9050808382526020820190506020840283018581111561046e5761046d610435565b5b835b8181101561049757806104838882610213565b845260208401935050602081019050610470565b5050509392505050565b5f82601f8301126104b5576104b4610406565b5b81356104c5848260208601610439565b91505092915050565b5f606082840312156104e3576104e2610376565b5b6104ed60606103e8565b90505f6104fc84828501610213565b5f83015250602061050f84828501610213565b602083015250604082013567ffffffffffffffff81111561053357610532610402565b5b61053f848285016104a1565b60408301525092915050565b5f5f60408385031215610561576105606101ec565b5b5f61056e85828601610213565b925050602083013567ffffffffffffffff81111561058f5761058e6101f0565b5b61059b858286016104ce565b915050925092905056fea2646970667358221220a9d6cfb5ff196c504dce43031a09ab08dd09ba8787f879f689fcd2f7b159d2df64736f6c634300081d0033",
}

// TransactionsStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionsStoreMetaData.ABI instead.
var TransactionsStoreABI = TransactionsStoreMetaData.ABI

// TransactionsStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransactionsStoreMetaData.Bin instead.
var TransactionsStoreBin = TransactionsStoreMetaData.Bin

// DeployTransactionsStore deploys a new Ethereum contract, binding an instance of TransactionsStore to it.
func DeployTransactionsStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransactionsStore, error) {
	parsed, err := TransactionsStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransactionsStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransactionsStore{TransactionsStoreCaller: TransactionsStoreCaller{contract: contract}, TransactionsStoreTransactor: TransactionsStoreTransactor{contract: contract}, TransactionsStoreFilterer: TransactionsStoreFilterer{contract: contract}}, nil
}

// TransactionsStore is an auto generated Go binding around an Ethereum contract.
type TransactionsStore struct {
	TransactionsStoreCaller     // Read-only binding to the contract
	TransactionsStoreTransactor // Write-only binding to the contract
	TransactionsStoreFilterer   // Log filterer for contract events
}

// TransactionsStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionsStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionsStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionsStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionsStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionsStoreSession struct {
	Contract     *TransactionsStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TransactionsStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionsStoreCallerSession struct {
	Contract *TransactionsStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TransactionsStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionsStoreTransactorSession struct {
	Contract     *TransactionsStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransactionsStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionsStoreRaw struct {
	Contract *TransactionsStore // Generic contract binding to access the raw methods on
}

// TransactionsStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionsStoreCallerRaw struct {
	Contract *TransactionsStoreCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionsStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionsStoreTransactorRaw struct {
	Contract *TransactionsStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionsStore creates a new instance of TransactionsStore, bound to a specific deployed contract.
func NewTransactionsStore(address common.Address, backend bind.ContractBackend) (*TransactionsStore, error) {
	contract, err := bindTransactionsStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransactionsStore{TransactionsStoreCaller: TransactionsStoreCaller{contract: contract}, TransactionsStoreTransactor: TransactionsStoreTransactor{contract: contract}, TransactionsStoreFilterer: TransactionsStoreFilterer{contract: contract}}, nil
}

// NewTransactionsStoreCaller creates a new read-only instance of TransactionsStore, bound to a specific deployed contract.
func NewTransactionsStoreCaller(address common.Address, caller bind.ContractCaller) (*TransactionsStoreCaller, error) {
	contract, err := bindTransactionsStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionsStoreCaller{contract: contract}, nil
}

// NewTransactionsStoreTransactor creates a new write-only instance of TransactionsStore, bound to a specific deployed contract.
func NewTransactionsStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionsStoreTransactor, error) {
	contract, err := bindTransactionsStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionsStoreTransactor{contract: contract}, nil
}

// NewTransactionsStoreFilterer creates a new log filterer instance of TransactionsStore, bound to a specific deployed contract.
func NewTransactionsStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionsStoreFilterer, error) {
	contract, err := bindTransactionsStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionsStoreFilterer{contract: contract}, nil
}

// bindTransactionsStore binds a generic wrapper to an already deployed contract.
func bindTransactionsStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransactionsStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionsStore *TransactionsStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionsStore.Contract.TransactionsStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionsStore *TransactionsStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionsStore.Contract.TransactionsStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionsStore *TransactionsStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionsStore.Contract.TransactionsStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionsStore *TransactionsStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionsStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionsStore *TransactionsStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionsStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionsStore *TransactionsStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionsStore.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x30c63439.
//
// Solidity: function GetAuction(uint256 id) view returns((uint256,uint256,uint256[]))
func (_TransactionsStore *TransactionsStoreCaller) GetAuction(opts *bind.CallOpts, id *big.Int) (TransactionmodelAuction, error) {
	var out []interface{}
	err := _TransactionsStore.contract.Call(opts, &out, "GetAuction", id)

	if err != nil {
		return *new(TransactionmodelAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(TransactionmodelAuction)).(*TransactionmodelAuction)

	return out0, err

}

// GetAuction is a free data retrieval call binding the contract method 0x30c63439.
//
// Solidity: function GetAuction(uint256 id) view returns((uint256,uint256,uint256[]))
func (_TransactionsStore *TransactionsStoreSession) GetAuction(id *big.Int) (TransactionmodelAuction, error) {
	return _TransactionsStore.Contract.GetAuction(&_TransactionsStore.CallOpts, id)
}

// GetAuction is a free data retrieval call binding the contract method 0x30c63439.
//
// Solidity: function GetAuction(uint256 id) view returns((uint256,uint256,uint256[]))
func (_TransactionsStore *TransactionsStoreCallerSession) GetAuction(id *big.Int) (TransactionmodelAuction, error) {
	return _TransactionsStore.Contract.GetAuction(&_TransactionsStore.CallOpts, id)
}

// SetAuction is a paid mutator transaction binding the contract method 0xf41c5ff7.
//
// Solidity: function SetAuction(uint256 id, (uint256,uint256,uint256[]) _auction) returns()
func (_TransactionsStore *TransactionsStoreTransactor) SetAuction(opts *bind.TransactOpts, id *big.Int, _auction TransactionmodelAuction) (*types.Transaction, error) {
	return _TransactionsStore.contract.Transact(opts, "SetAuction", id, _auction)
}

// SetAuction is a paid mutator transaction binding the contract method 0xf41c5ff7.
//
// Solidity: function SetAuction(uint256 id, (uint256,uint256,uint256[]) _auction) returns()
func (_TransactionsStore *TransactionsStoreSession) SetAuction(id *big.Int, _auction TransactionmodelAuction) (*types.Transaction, error) {
	return _TransactionsStore.Contract.SetAuction(&_TransactionsStore.TransactOpts, id, _auction)
}

// SetAuction is a paid mutator transaction binding the contract method 0xf41c5ff7.
//
// Solidity: function SetAuction(uint256 id, (uint256,uint256,uint256[]) _auction) returns()
func (_TransactionsStore *TransactionsStoreTransactorSession) SetAuction(id *big.Int, _auction TransactionmodelAuction) (*types.Transaction, error) {
	return _TransactionsStore.Contract.SetAuction(&_TransactionsStore.TransactOpts, id, _auction)
}

// OrgmodelMetaData contains all meta data concerning the Orgmodel contract.
var OrgmodelMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212201e9231772cebf1a5a4db867b9ed7bbc35e0bb59d2df37e5d70614d233f65349f64736f6c634300081d0033",
}

// OrgmodelABI is the input ABI used to generate the binding from.
// Deprecated: Use OrgmodelMetaData.ABI instead.
var OrgmodelABI = OrgmodelMetaData.ABI

// OrgmodelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrgmodelMetaData.Bin instead.
var OrgmodelBin = OrgmodelMetaData.Bin

// DeployOrgmodel deploys a new Ethereum contract, binding an instance of Orgmodel to it.
func DeployOrgmodel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orgmodel, error) {
	parsed, err := OrgmodelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrgmodelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orgmodel{OrgmodelCaller: OrgmodelCaller{contract: contract}, OrgmodelTransactor: OrgmodelTransactor{contract: contract}, OrgmodelFilterer: OrgmodelFilterer{contract: contract}}, nil
}

// Orgmodel is an auto generated Go binding around an Ethereum contract.
type Orgmodel struct {
	OrgmodelCaller     // Read-only binding to the contract
	OrgmodelTransactor // Write-only binding to the contract
	OrgmodelFilterer   // Log filterer for contract events
}

// OrgmodelCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgmodelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgmodelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgmodelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgmodelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgmodelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgmodelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgmodelSession struct {
	Contract     *Orgmodel         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgmodelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgmodelCallerSession struct {
	Contract *OrgmodelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OrgmodelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgmodelTransactorSession struct {
	Contract     *OrgmodelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OrgmodelRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgmodelRaw struct {
	Contract *Orgmodel // Generic contract binding to access the raw methods on
}

// OrgmodelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgmodelCallerRaw struct {
	Contract *OrgmodelCaller // Generic read-only contract binding to access the raw methods on
}

// OrgmodelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgmodelTransactorRaw struct {
	Contract *OrgmodelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgmodel creates a new instance of Orgmodel, bound to a specific deployed contract.
func NewOrgmodel(address common.Address, backend bind.ContractBackend) (*Orgmodel, error) {
	contract, err := bindOrgmodel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orgmodel{OrgmodelCaller: OrgmodelCaller{contract: contract}, OrgmodelTransactor: OrgmodelTransactor{contract: contract}, OrgmodelFilterer: OrgmodelFilterer{contract: contract}}, nil
}

// NewOrgmodelCaller creates a new read-only instance of Orgmodel, bound to a specific deployed contract.
func NewOrgmodelCaller(address common.Address, caller bind.ContractCaller) (*OrgmodelCaller, error) {
	contract, err := bindOrgmodel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgmodelCaller{contract: contract}, nil
}

// NewOrgmodelTransactor creates a new write-only instance of Orgmodel, bound to a specific deployed contract.
func NewOrgmodelTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgmodelTransactor, error) {
	contract, err := bindOrgmodel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgmodelTransactor{contract: contract}, nil
}

// NewOrgmodelFilterer creates a new log filterer instance of Orgmodel, bound to a specific deployed contract.
func NewOrgmodelFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgmodelFilterer, error) {
	contract, err := bindOrgmodel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgmodelFilterer{contract: contract}, nil
}

// bindOrgmodel binds a generic wrapper to an already deployed contract.
func bindOrgmodel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrgmodelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orgmodel *OrgmodelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orgmodel.Contract.OrgmodelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orgmodel *OrgmodelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orgmodel.Contract.OrgmodelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orgmodel *OrgmodelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orgmodel.Contract.OrgmodelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orgmodel *OrgmodelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orgmodel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orgmodel *OrgmodelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orgmodel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orgmodel *OrgmodelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orgmodel.Contract.contract.Transact(opts, method, params...)
}

// RequestmodelMetaData contains all meta data concerning the Requestmodel contract.
var RequestmodelMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212207d056acf6ffea1869581ead5a1da6ee60eae6d2334795a7b5dc34b5d37bef23664736f6c634300081d0033",
}

// RequestmodelABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestmodelMetaData.ABI instead.
var RequestmodelABI = RequestmodelMetaData.ABI

// RequestmodelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestmodelMetaData.Bin instead.
var RequestmodelBin = RequestmodelMetaData.Bin

// DeployRequestmodel deploys a new Ethereum contract, binding an instance of Requestmodel to it.
func DeployRequestmodel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Requestmodel, error) {
	parsed, err := RequestmodelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestmodelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Requestmodel{RequestmodelCaller: RequestmodelCaller{contract: contract}, RequestmodelTransactor: RequestmodelTransactor{contract: contract}, RequestmodelFilterer: RequestmodelFilterer{contract: contract}}, nil
}

// Requestmodel is an auto generated Go binding around an Ethereum contract.
type Requestmodel struct {
	RequestmodelCaller     // Read-only binding to the contract
	RequestmodelTransactor // Write-only binding to the contract
	RequestmodelFilterer   // Log filterer for contract events
}

// RequestmodelCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestmodelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestmodelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestmodelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestmodelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestmodelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestmodelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestmodelSession struct {
	Contract     *Requestmodel     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestmodelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestmodelCallerSession struct {
	Contract *RequestmodelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RequestmodelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestmodelTransactorSession struct {
	Contract     *RequestmodelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RequestmodelRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestmodelRaw struct {
	Contract *Requestmodel // Generic contract binding to access the raw methods on
}

// RequestmodelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestmodelCallerRaw struct {
	Contract *RequestmodelCaller // Generic read-only contract binding to access the raw methods on
}

// RequestmodelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestmodelTransactorRaw struct {
	Contract *RequestmodelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestmodel creates a new instance of Requestmodel, bound to a specific deployed contract.
func NewRequestmodel(address common.Address, backend bind.ContractBackend) (*Requestmodel, error) {
	contract, err := bindRequestmodel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Requestmodel{RequestmodelCaller: RequestmodelCaller{contract: contract}, RequestmodelTransactor: RequestmodelTransactor{contract: contract}, RequestmodelFilterer: RequestmodelFilterer{contract: contract}}, nil
}

// NewRequestmodelCaller creates a new read-only instance of Requestmodel, bound to a specific deployed contract.
func NewRequestmodelCaller(address common.Address, caller bind.ContractCaller) (*RequestmodelCaller, error) {
	contract, err := bindRequestmodel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestmodelCaller{contract: contract}, nil
}

// NewRequestmodelTransactor creates a new write-only instance of Requestmodel, bound to a specific deployed contract.
func NewRequestmodelTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestmodelTransactor, error) {
	contract, err := bindRequestmodel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestmodelTransactor{contract: contract}, nil
}

// NewRequestmodelFilterer creates a new log filterer instance of Requestmodel, bound to a specific deployed contract.
func NewRequestmodelFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestmodelFilterer, error) {
	contract, err := bindRequestmodel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestmodelFilterer{contract: contract}, nil
}

// bindRequestmodel binds a generic wrapper to an already deployed contract.
func bindRequestmodel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestmodelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Requestmodel *RequestmodelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Requestmodel.Contract.RequestmodelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Requestmodel *RequestmodelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Requestmodel.Contract.RequestmodelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Requestmodel *RequestmodelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Requestmodel.Contract.RequestmodelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Requestmodel *RequestmodelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Requestmodel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Requestmodel *RequestmodelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Requestmodel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Requestmodel *RequestmodelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Requestmodel.Contract.contract.Transact(opts, method, params...)
}

// RequeststoreMetaData contains all meta data concerning the Requeststore contract.
var RequeststoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumrequestmodel.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structrequestmodel.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumrequestmodel.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structrequestmodel.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"SaveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610a8e8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80639b37031714610038578063c997fe8b14610068575b5f5ffd5b610052600480360381019061004d9190610344565b610084565b60405161005f9190610500565b60405180910390f35b610082600480360381019061007d9190610732565b6101d7565b005b61008c6102b3565b5f5f8381526020019081526020015f206040518060800160405290815f82015481526020016001820180546100c0906107b9565b80601f01602080910402602001604051908101604052809291908181526020018280546100ec906107b9565b80156101375780601f1061010e57610100808354040283529160200191610137565b820191905f5260205f20905b81548152906001019060200180831161011a57829003601f168201915b50505050508152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff1660048111156101ba576101b961042d565b5b60048111156101cc576101cb61042d565b5b815250509050919050565b5f816060019060048111156101ef576101ee61042d565b5b908160048111156102035761020261042d565b5b81525050805f5f8481526020019081526020015f205f820151815f015560208201518160010190816102359190610989565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff021916908360048111156102a7576102a661042d565b5b02179055509050505050565b60405180608001604052805f8152602001606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f60048111156102fa576102f961042d565b5b81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61032381610311565b811461032d575f5ffd5b50565b5f8135905061033e8161031a565b92915050565b5f6020828403121561035957610358610309565b5b5f61036684828501610330565b91505092915050565b61037881610311565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6103c08261037e565b6103ca8185610388565b93506103da818560208601610398565b6103e3816103a6565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610417826103ee565b9050919050565b6104278161040d565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061046b5761046a61042d565b5b50565b5f81905061047b8261045a565b919050565b5f61048a8261046e565b9050919050565b61049a81610480565b82525050565b5f608083015f8301516104b55f86018261036f565b50602083015184820360208601526104cd82826103b6565b91505060408301516104e2604086018261041e565b5060608301516104f56060860182610491565b508091505092915050565b5f6020820190508181035f83015261051881846104a0565b905092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61055a826103a6565b810181811067ffffffffffffffff8211171561057957610578610524565b5b80604052505050565b5f61058b610300565b90506105978282610551565b919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f67ffffffffffffffff8211156105c2576105c1610524565b5b6105cb826103a6565b9050602081019050919050565b828183375f83830152505050565b5f6105f86105f3846105a8565b610582565b905082815260208101848484011115610614576106136105a4565b5b61061f8482856105d8565b509392505050565b5f82601f83011261063b5761063a6105a0565b5b813561064b8482602086016105e6565b91505092915050565b61065d8161040d565b8114610667575f5ffd5b50565b5f8135905061067881610654565b92915050565b6005811061068a575f5ffd5b50565b5f8135905061069b8161067e565b92915050565b5f608082840312156106b6576106b5610520565b5b6106c06080610582565b90505f6106cf84828501610330565b5f83015250602082013567ffffffffffffffff8111156106f2576106f161059c565b5b6106fe84828501610627565b60208301525060406107128482850161066a565b60408301525060606107268482850161068d565b60608301525092915050565b5f5f6040838503121561074857610747610309565b5b5f61075585828601610330565b925050602083013567ffffffffffffffff8111156107765761077561030d565b5b610782858286016106a1565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806107d057607f821691505b6020821081036107e3576107e261078c565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108457fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261080a565b61084f868361080a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61088a61088561088084610311565b610867565b610311565b9050919050565b5f819050919050565b6108a383610870565b6108b76108af82610891565b848454610816565b825550505050565b5f5f905090565b6108ce6108bf565b6108d981848461089a565b505050565b5b818110156108fc576108f15f826108c6565b6001810190506108df565b5050565b601f82111561094157610912816107e9565b61091b846107fb565b8101602085101561092a578190505b61093e610936856107fb565b8301826108de565b50505b505050565b5f82821c905092915050565b5f6109615f1984600802610946565b1980831691505092915050565b5f6109798383610952565b9150826002028217905092915050565b6109928261037e565b67ffffffffffffffff8111156109ab576109aa610524565b5b6109b582546107b9565b6109c0828285610900565b5f60209050601f8311600181146109f1575f84156109df578287015190505b6109e9858261096e565b865550610a50565b601f1984166109ff866107e9565b5f5b82811015610a2657848901518255600182019150602085019450602081019050610a01565b86831015610a435784890151610a3f601f891682610952565b8355505b6001600288020188555050505b50505050505056fea264697066735822122098506fd5725ca1527056c86563e28eac2c278624241e4dbf67b7de8f70a19a0f64736f6c634300081d0033",
}

// RequeststoreABI is the input ABI used to generate the binding from.
// Deprecated: Use RequeststoreMetaData.ABI instead.
var RequeststoreABI = RequeststoreMetaData.ABI

// RequeststoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequeststoreMetaData.Bin instead.
var RequeststoreBin = RequeststoreMetaData.Bin

// DeployRequeststore deploys a new Ethereum contract, binding an instance of Requeststore to it.
func DeployRequeststore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Requeststore, error) {
	parsed, err := RequeststoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequeststoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Requeststore{RequeststoreCaller: RequeststoreCaller{contract: contract}, RequeststoreTransactor: RequeststoreTransactor{contract: contract}, RequeststoreFilterer: RequeststoreFilterer{contract: contract}}, nil
}

// Requeststore is an auto generated Go binding around an Ethereum contract.
type Requeststore struct {
	RequeststoreCaller     // Read-only binding to the contract
	RequeststoreTransactor // Write-only binding to the contract
	RequeststoreFilterer   // Log filterer for contract events
}

// RequeststoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequeststoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequeststoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequeststoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequeststoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequeststoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequeststoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequeststoreSession struct {
	Contract     *Requeststore     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequeststoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequeststoreCallerSession struct {
	Contract *RequeststoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RequeststoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequeststoreTransactorSession struct {
	Contract     *RequeststoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RequeststoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequeststoreRaw struct {
	Contract *Requeststore // Generic contract binding to access the raw methods on
}

// RequeststoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequeststoreCallerRaw struct {
	Contract *RequeststoreCaller // Generic read-only contract binding to access the raw methods on
}

// RequeststoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequeststoreTransactorRaw struct {
	Contract *RequeststoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequeststore creates a new instance of Requeststore, bound to a specific deployed contract.
func NewRequeststore(address common.Address, backend bind.ContractBackend) (*Requeststore, error) {
	contract, err := bindRequeststore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Requeststore{RequeststoreCaller: RequeststoreCaller{contract: contract}, RequeststoreTransactor: RequeststoreTransactor{contract: contract}, RequeststoreFilterer: RequeststoreFilterer{contract: contract}}, nil
}

// NewRequeststoreCaller creates a new read-only instance of Requeststore, bound to a specific deployed contract.
func NewRequeststoreCaller(address common.Address, caller bind.ContractCaller) (*RequeststoreCaller, error) {
	contract, err := bindRequeststore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequeststoreCaller{contract: contract}, nil
}

// NewRequeststoreTransactor creates a new write-only instance of Requeststore, bound to a specific deployed contract.
func NewRequeststoreTransactor(address common.Address, transactor bind.ContractTransactor) (*RequeststoreTransactor, error) {
	contract, err := bindRequeststore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequeststoreTransactor{contract: contract}, nil
}

// NewRequeststoreFilterer creates a new log filterer instance of Requeststore, bound to a specific deployed contract.
func NewRequeststoreFilterer(address common.Address, filterer bind.ContractFilterer) (*RequeststoreFilterer, error) {
	contract, err := bindRequeststore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequeststoreFilterer{contract: contract}, nil
}

// bindRequeststore binds a generic wrapper to an already deployed contract.
func bindRequeststore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequeststoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Requeststore *RequeststoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Requeststore.Contract.RequeststoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Requeststore *RequeststoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Requeststore.Contract.RequeststoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Requeststore *RequeststoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Requeststore.Contract.RequeststoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Requeststore *RequeststoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Requeststore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Requeststore *RequeststoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Requeststore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Requeststore *RequeststoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Requeststore.Contract.contract.Transact(opts, method, params...)
}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_Requeststore *RequeststoreCaller) GetRequest(opts *bind.CallOpts, id *big.Int) (RequestmodelRequest, error) {
	var out []interface{}
	err := _Requeststore.contract.Call(opts, &out, "GetRequest", id)

	if err != nil {
		return *new(RequestmodelRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(RequestmodelRequest)).(*RequestmodelRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_Requeststore *RequeststoreSession) GetRequest(id *big.Int) (RequestmodelRequest, error) {
	return _Requeststore.Contract.GetRequest(&_Requeststore.CallOpts, id)
}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(uint256 id) view returns((uint256,string,address,uint8))
func (_Requeststore *RequeststoreCallerSession) GetRequest(id *big.Int) (RequestmodelRequest, error) {
	return _Requeststore.Contract.GetRequest(&_Requeststore.CallOpts, id)
}

// SaveRequest is a paid mutator transaction binding the contract method 0xc997fe8b.
//
// Solidity: function SaveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_Requeststore *RequeststoreTransactor) SaveRequest(opts *bind.TransactOpts, id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _Requeststore.contract.Transact(opts, "SaveRequest", id, _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0xc997fe8b.
//
// Solidity: function SaveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_Requeststore *RequeststoreSession) SaveRequest(id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _Requeststore.Contract.SaveRequest(&_Requeststore.TransactOpts, id, _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0xc997fe8b.
//
// Solidity: function SaveRequest(uint256 id, (uint256,string,address,uint8) _request) returns()
func (_Requeststore *RequeststoreTransactorSession) SaveRequest(id *big.Int, _request RequestmodelRequest) (*types.Transaction, error) {
	return _Requeststore.Contract.SaveRequest(&_Requeststore.TransactOpts, id, _request)
}

// TokenmodelMetaData contains all meta data concerning the Tokenmodel contract.
var TokenmodelMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212209dba1486a1ea109b03ed16134e6d2b874f29da5da7c7685744e06bab95fdd63c64736f6c634300081d0033",
}

// TokenmodelABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenmodelMetaData.ABI instead.
var TokenmodelABI = TokenmodelMetaData.ABI

// TokenmodelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenmodelMetaData.Bin instead.
var TokenmodelBin = TokenmodelMetaData.Bin

// DeployTokenmodel deploys a new Ethereum contract, binding an instance of Tokenmodel to it.
func DeployTokenmodel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenmodel, error) {
	parsed, err := TokenmodelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenmodelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenmodel{TokenmodelCaller: TokenmodelCaller{contract: contract}, TokenmodelTransactor: TokenmodelTransactor{contract: contract}, TokenmodelFilterer: TokenmodelFilterer{contract: contract}}, nil
}

// Tokenmodel is an auto generated Go binding around an Ethereum contract.
type Tokenmodel struct {
	TokenmodelCaller     // Read-only binding to the contract
	TokenmodelTransactor // Write-only binding to the contract
	TokenmodelFilterer   // Log filterer for contract events
}

// TokenmodelCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenmodelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmodelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenmodelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmodelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenmodelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmodelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenmodelSession struct {
	Contract     *Tokenmodel       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenmodelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenmodelCallerSession struct {
	Contract *TokenmodelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenmodelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenmodelTransactorSession struct {
	Contract     *TokenmodelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenmodelRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenmodelRaw struct {
	Contract *Tokenmodel // Generic contract binding to access the raw methods on
}

// TokenmodelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenmodelCallerRaw struct {
	Contract *TokenmodelCaller // Generic read-only contract binding to access the raw methods on
}

// TokenmodelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenmodelTransactorRaw struct {
	Contract *TokenmodelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenmodel creates a new instance of Tokenmodel, bound to a specific deployed contract.
func NewTokenmodel(address common.Address, backend bind.ContractBackend) (*Tokenmodel, error) {
	contract, err := bindTokenmodel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenmodel{TokenmodelCaller: TokenmodelCaller{contract: contract}, TokenmodelTransactor: TokenmodelTransactor{contract: contract}, TokenmodelFilterer: TokenmodelFilterer{contract: contract}}, nil
}

// NewTokenmodelCaller creates a new read-only instance of Tokenmodel, bound to a specific deployed contract.
func NewTokenmodelCaller(address common.Address, caller bind.ContractCaller) (*TokenmodelCaller, error) {
	contract, err := bindTokenmodel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmodelCaller{contract: contract}, nil
}

// NewTokenmodelTransactor creates a new write-only instance of Tokenmodel, bound to a specific deployed contract.
func NewTokenmodelTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenmodelTransactor, error) {
	contract, err := bindTokenmodel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmodelTransactor{contract: contract}, nil
}

// NewTokenmodelFilterer creates a new log filterer instance of Tokenmodel, bound to a specific deployed contract.
func NewTokenmodelFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenmodelFilterer, error) {
	contract, err := bindTokenmodel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenmodelFilterer{contract: contract}, nil
}

// bindTokenmodel binds a generic wrapper to an already deployed contract.
func bindTokenmodel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenmodelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenmodel *TokenmodelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenmodel.Contract.TokenmodelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenmodel *TokenmodelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenmodel.Contract.TokenmodelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenmodel *TokenmodelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenmodel.Contract.TokenmodelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenmodel *TokenmodelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenmodel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenmodel *TokenmodelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenmodel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenmodel *TokenmodelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenmodel.Contract.contract.Transact(opts, method, params...)
}

// TokenstoreMetaData contains all meta data concerning the Tokenstore contract.
var TokenstoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GetNewTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumtokenmodel.TokenStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structtokenmodel.Token\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storageId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashdata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumtokenmodel.TokenStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structtokenmodel.Token\",\"name\":\"_token\",\"type\":\"tuple\"}],\"name\":\"SetToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumtokenmodel.TokenStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"UpdateTokenStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610ca58061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806340b93af114610059578063512b274d14610075578063abb1a50314610091578063aef15d37146100ad578063afcb7a85146100cb575b5f5ffd5b610073600480360381019061006e919061049d565b6100fb565b005b61008f600480360381019061008a91906106d3565b61014e565b005b6100ab60048036038101906100a6919061072d565b6101fa565b005b6100b5610245565b6040516100c2919061077a565b60405180910390f35b6100e560048036038101906100e09190610793565b61025f565b6040516100f2919061090f565b60405180910390f35b5f6101058361025f565b905081816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050610149838261014e565b505050565b805f5f8481526020019081526020015f205f820151815f0155602082015181600101908161017c9190610b2c565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff021916908360048111156101ee576101ed61083c565b5b02179055509050505050565b5f6102048361025f565b9050818160600190600481111561021e5761021d61083c565b5b908160048111156102325761023161083c565b5b81525050610240838261014e565b505050565b5f60015f815461025490610c28565b919050819055905090565b6102676103b2565b5f5f8381526020019081526020015f206040518060800160405290815f820154815260200160018201805461029b9061095c565b80601f01602080910402602001604051908101604052809291908181526020018280546102c79061095c565b80156103125780601f106102e957610100808354040283529160200191610312565b820191905f5260205f20905b8154815290600101906020018083116102f557829003601f168201915b50505050508152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff1660048111156103955761039461083c565b5b60048111156103a7576103a661083c565b5b815250509050919050565b60405180608001604052805f8152602001606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f60048111156103f9576103f861083c565b5b81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61042281610410565b811461042c575f5ffd5b50565b5f8135905061043d81610419565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61046c82610443565b9050919050565b61047c81610462565b8114610486575f5ffd5b50565b5f8135905061049781610473565b92915050565b5f5f604083850312156104b3576104b2610408565b5b5f6104c08582860161042f565b92505060206104d185828601610489565b9150509250929050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610525826104df565b810181811067ffffffffffffffff82111715610544576105436104ef565b5b80604052505050565b5f6105566103ff565b9050610562828261051c565b919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82111561058d5761058c6104ef565b5b610596826104df565b9050602081019050919050565b828183375f83830152505050565b5f6105c36105be84610573565b61054d565b9050828152602081018484840111156105df576105de61056f565b5b6105ea8482856105a3565b509392505050565b5f82601f8301126106065761060561056b565b5b81356106168482602086016105b1565b91505092915050565b6005811061062b575f5ffd5b50565b5f8135905061063c8161061f565b92915050565b5f60808284031215610657576106566104db565b5b610661608061054d565b90505f6106708482850161042f565b5f83015250602082013567ffffffffffffffff81111561069357610692610567565b5b61069f848285016105f2565b60208301525060406106b384828501610489565b60408301525060606106c78482850161062e565b60608301525092915050565b5f5f604083850312156106e9576106e8610408565b5b5f6106f68582860161042f565b925050602083013567ffffffffffffffff8111156107175761071661040c565b5b61072385828601610642565b9150509250929050565b5f5f6040838503121561074357610742610408565b5b5f6107508582860161042f565b92505060206107618582860161062e565b9150509250929050565b61077481610410565b82525050565b5f60208201905061078d5f83018461076b565b92915050565b5f602082840312156107a8576107a7610408565b5b5f6107b58482850161042f565b91505092915050565b6107c781610410565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6107ff826107cd565b61080981856107d7565b93506108198185602086016107e7565b610822816104df565b840191505092915050565b61083681610462565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061087a5761087961083c565b5b50565b5f81905061088a82610869565b919050565b5f6108998261087d565b9050919050565b6108a98161088f565b82525050565b5f608083015f8301516108c45f8601826107be565b50602083015184820360208601526108dc82826107f5565b91505060408301516108f1604086018261082d565b50606083015161090460608601826108a0565b508091505092915050565b5f6020820190508181035f83015261092781846108af565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061097357607f821691505b6020821081036109865761098561092f565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109e87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109ad565b6109f286836109ad565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610a2d610a28610a2384610410565b610a0a565b610410565b9050919050565b5f819050919050565b610a4683610a13565b610a5a610a5282610a34565b8484546109b9565b825550505050565b5f5f905090565b610a71610a62565b610a7c818484610a3d565b505050565b5b81811015610a9f57610a945f82610a69565b600181019050610a82565b5050565b601f821115610ae457610ab58161098c565b610abe8461099e565b81016020851015610acd578190505b610ae1610ad98561099e565b830182610a81565b50505b505050565b5f82821c905092915050565b5f610b045f1984600802610ae9565b1980831691505092915050565b5f610b1c8383610af5565b9150826002028217905092915050565b610b35826107cd565b67ffffffffffffffff811115610b4e57610b4d6104ef565b5b610b58825461095c565b610b63828285610aa3565b5f60209050601f831160018114610b94575f8415610b82578287015190505b610b8c8582610b11565b865550610bf3565b601f198416610ba28661098c565b5f5b82811015610bc957848901518255600182019150602085019450602081019050610ba4565b86831015610be65784890151610be2601f891682610af5565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c3282610410565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c6457610c63610bfb565b5b60018201905091905056fea26469706673582212204b56864057b6e59d970cc43200635240856980f703687ecd7f89aa50803ab6ab64736f6c634300081d0033",
}

// TokenstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenstoreMetaData.ABI instead.
var TokenstoreABI = TokenstoreMetaData.ABI

// TokenstoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenstoreMetaData.Bin instead.
var TokenstoreBin = TokenstoreMetaData.Bin

// DeployTokenstore deploys a new Ethereum contract, binding an instance of Tokenstore to it.
func DeployTokenstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenstore, error) {
	parsed, err := TokenstoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenstoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenstore{TokenstoreCaller: TokenstoreCaller{contract: contract}, TokenstoreTransactor: TokenstoreTransactor{contract: contract}, TokenstoreFilterer: TokenstoreFilterer{contract: contract}}, nil
}

// Tokenstore is an auto generated Go binding around an Ethereum contract.
type Tokenstore struct {
	TokenstoreCaller     // Read-only binding to the contract
	TokenstoreTransactor // Write-only binding to the contract
	TokenstoreFilterer   // Log filterer for contract events
}

// TokenstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenstoreSession struct {
	Contract     *Tokenstore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenstoreCallerSession struct {
	Contract *TokenstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenstoreTransactorSession struct {
	Contract     *TokenstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenstoreRaw struct {
	Contract *Tokenstore // Generic contract binding to access the raw methods on
}

// TokenstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenstoreCallerRaw struct {
	Contract *TokenstoreCaller // Generic read-only contract binding to access the raw methods on
}

// TokenstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenstoreTransactorRaw struct {
	Contract *TokenstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenstore creates a new instance of Tokenstore, bound to a specific deployed contract.
func NewTokenstore(address common.Address, backend bind.ContractBackend) (*Tokenstore, error) {
	contract, err := bindTokenstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenstore{TokenstoreCaller: TokenstoreCaller{contract: contract}, TokenstoreTransactor: TokenstoreTransactor{contract: contract}, TokenstoreFilterer: TokenstoreFilterer{contract: contract}}, nil
}

// NewTokenstoreCaller creates a new read-only instance of Tokenstore, bound to a specific deployed contract.
func NewTokenstoreCaller(address common.Address, caller bind.ContractCaller) (*TokenstoreCaller, error) {
	contract, err := bindTokenstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenstoreCaller{contract: contract}, nil
}

// NewTokenstoreTransactor creates a new write-only instance of Tokenstore, bound to a specific deployed contract.
func NewTokenstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenstoreTransactor, error) {
	contract, err := bindTokenstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenstoreTransactor{contract: contract}, nil
}

// NewTokenstoreFilterer creates a new log filterer instance of Tokenstore, bound to a specific deployed contract.
func NewTokenstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenstoreFilterer, error) {
	contract, err := bindTokenstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenstoreFilterer{contract: contract}, nil
}

// bindTokenstore binds a generic wrapper to an already deployed contract.
func bindTokenstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenstoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenstore *TokenstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenstore.Contract.TokenstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenstore *TokenstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenstore.Contract.TokenstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenstore *TokenstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenstore.Contract.TokenstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenstore *TokenstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenstore *TokenstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenstore *TokenstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenstore.Contract.contract.Transact(opts, method, params...)
}

// GetToken is a free data retrieval call binding the contract method 0xafcb7a85.
//
// Solidity: function GetToken(uint256 id) view returns((uint256,string,address,uint8))
func (_Tokenstore *TokenstoreCaller) GetToken(opts *bind.CallOpts, id *big.Int) (TokenmodelToken, error) {
	var out []interface{}
	err := _Tokenstore.contract.Call(opts, &out, "GetToken", id)

	if err != nil {
		return *new(TokenmodelToken), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenmodelToken)).(*TokenmodelToken)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0xafcb7a85.
//
// Solidity: function GetToken(uint256 id) view returns((uint256,string,address,uint8))
func (_Tokenstore *TokenstoreSession) GetToken(id *big.Int) (TokenmodelToken, error) {
	return _Tokenstore.Contract.GetToken(&_Tokenstore.CallOpts, id)
}

// GetToken is a free data retrieval call binding the contract method 0xafcb7a85.
//
// Solidity: function GetToken(uint256 id) view returns((uint256,string,address,uint8))
func (_Tokenstore *TokenstoreCallerSession) GetToken(id *big.Int) (TokenmodelToken, error) {
	return _Tokenstore.Contract.GetToken(&_Tokenstore.CallOpts, id)
}

// GetNewTokenId is a paid mutator transaction binding the contract method 0xaef15d37.
//
// Solidity: function GetNewTokenId() returns(uint256)
func (_Tokenstore *TokenstoreTransactor) GetNewTokenId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenstore.contract.Transact(opts, "GetNewTokenId")
}

// GetNewTokenId is a paid mutator transaction binding the contract method 0xaef15d37.
//
// Solidity: function GetNewTokenId() returns(uint256)
func (_Tokenstore *TokenstoreSession) GetNewTokenId() (*types.Transaction, error) {
	return _Tokenstore.Contract.GetNewTokenId(&_Tokenstore.TransactOpts)
}

// GetNewTokenId is a paid mutator transaction binding the contract method 0xaef15d37.
//
// Solidity: function GetNewTokenId() returns(uint256)
func (_Tokenstore *TokenstoreTransactorSession) GetNewTokenId() (*types.Transaction, error) {
	return _Tokenstore.Contract.GetNewTokenId(&_Tokenstore.TransactOpts)
}

// SetToken is a paid mutator transaction binding the contract method 0x512b274d.
//
// Solidity: function SetToken(uint256 id, (uint256,string,address,uint8) _token) returns()
func (_Tokenstore *TokenstoreTransactor) SetToken(opts *bind.TransactOpts, id *big.Int, _token TokenmodelToken) (*types.Transaction, error) {
	return _Tokenstore.contract.Transact(opts, "SetToken", id, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x512b274d.
//
// Solidity: function SetToken(uint256 id, (uint256,string,address,uint8) _token) returns()
func (_Tokenstore *TokenstoreSession) SetToken(id *big.Int, _token TokenmodelToken) (*types.Transaction, error) {
	return _Tokenstore.Contract.SetToken(&_Tokenstore.TransactOpts, id, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x512b274d.
//
// Solidity: function SetToken(uint256 id, (uint256,string,address,uint8) _token) returns()
func (_Tokenstore *TokenstoreTransactorSession) SetToken(id *big.Int, _token TokenmodelToken) (*types.Transaction, error) {
	return _Tokenstore.Contract.SetToken(&_Tokenstore.TransactOpts, id, _token)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x40b93af1.
//
// Solidity: function UpdateOwner(uint256 id, address owner) returns()
func (_Tokenstore *TokenstoreTransactor) UpdateOwner(opts *bind.TransactOpts, id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Tokenstore.contract.Transact(opts, "UpdateOwner", id, owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x40b93af1.
//
// Solidity: function UpdateOwner(uint256 id, address owner) returns()
func (_Tokenstore *TokenstoreSession) UpdateOwner(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Tokenstore.Contract.UpdateOwner(&_Tokenstore.TransactOpts, id, owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x40b93af1.
//
// Solidity: function UpdateOwner(uint256 id, address owner) returns()
func (_Tokenstore *TokenstoreTransactorSession) UpdateOwner(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Tokenstore.Contract.UpdateOwner(&_Tokenstore.TransactOpts, id, owner)
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0xabb1a503.
//
// Solidity: function UpdateTokenStatus(uint256 id, uint8 _status) returns()
func (_Tokenstore *TokenstoreTransactor) UpdateTokenStatus(opts *bind.TransactOpts, id *big.Int, _status uint8) (*types.Transaction, error) {
	return _Tokenstore.contract.Transact(opts, "UpdateTokenStatus", id, _status)
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0xabb1a503.
//
// Solidity: function UpdateTokenStatus(uint256 id, uint8 _status) returns()
func (_Tokenstore *TokenstoreSession) UpdateTokenStatus(id *big.Int, _status uint8) (*types.Transaction, error) {
	return _Tokenstore.Contract.UpdateTokenStatus(&_Tokenstore.TransactOpts, id, _status)
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0xabb1a503.
//
// Solidity: function UpdateTokenStatus(uint256 id, uint8 _status) returns()
func (_Tokenstore *TokenstoreTransactorSession) UpdateTokenStatus(id *big.Int, _status uint8) (*types.Transaction, error) {
	return _Tokenstore.Contract.UpdateTokenStatus(&_Tokenstore.TransactOpts, id, _status)
}

// TransactionmodelMetaData contains all meta data concerning the Transactionmodel contract.
var TransactionmodelMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220bfdd8324ad625629abf67d62e968133f2315fd1b2a37d10af5a9ef1c11420c7564736f6c634300081d0033",
}

// TransactionmodelABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionmodelMetaData.ABI instead.
var TransactionmodelABI = TransactionmodelMetaData.ABI

// TransactionmodelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransactionmodelMetaData.Bin instead.
var TransactionmodelBin = TransactionmodelMetaData.Bin

// DeployTransactionmodel deploys a new Ethereum contract, binding an instance of Transactionmodel to it.
func DeployTransactionmodel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transactionmodel, error) {
	parsed, err := TransactionmodelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransactionmodelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transactionmodel{TransactionmodelCaller: TransactionmodelCaller{contract: contract}, TransactionmodelTransactor: TransactionmodelTransactor{contract: contract}, TransactionmodelFilterer: TransactionmodelFilterer{contract: contract}}, nil
}

// Transactionmodel is an auto generated Go binding around an Ethereum contract.
type Transactionmodel struct {
	TransactionmodelCaller     // Read-only binding to the contract
	TransactionmodelTransactor // Write-only binding to the contract
	TransactionmodelFilterer   // Log filterer for contract events
}

// TransactionmodelCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionmodelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionmodelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionmodelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionmodelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionmodelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionmodelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionmodelSession struct {
	Contract     *Transactionmodel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionmodelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionmodelCallerSession struct {
	Contract *TransactionmodelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TransactionmodelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionmodelTransactorSession struct {
	Contract     *TransactionmodelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TransactionmodelRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionmodelRaw struct {
	Contract *Transactionmodel // Generic contract binding to access the raw methods on
}

// TransactionmodelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionmodelCallerRaw struct {
	Contract *TransactionmodelCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionmodelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionmodelTransactorRaw struct {
	Contract *TransactionmodelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionmodel creates a new instance of Transactionmodel, bound to a specific deployed contract.
func NewTransactionmodel(address common.Address, backend bind.ContractBackend) (*Transactionmodel, error) {
	contract, err := bindTransactionmodel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transactionmodel{TransactionmodelCaller: TransactionmodelCaller{contract: contract}, TransactionmodelTransactor: TransactionmodelTransactor{contract: contract}, TransactionmodelFilterer: TransactionmodelFilterer{contract: contract}}, nil
}

// NewTransactionmodelCaller creates a new read-only instance of Transactionmodel, bound to a specific deployed contract.
func NewTransactionmodelCaller(address common.Address, caller bind.ContractCaller) (*TransactionmodelCaller, error) {
	contract, err := bindTransactionmodel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionmodelCaller{contract: contract}, nil
}

// NewTransactionmodelTransactor creates a new write-only instance of Transactionmodel, bound to a specific deployed contract.
func NewTransactionmodelTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionmodelTransactor, error) {
	contract, err := bindTransactionmodel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionmodelTransactor{contract: contract}, nil
}

// NewTransactionmodelFilterer creates a new log filterer instance of Transactionmodel, bound to a specific deployed contract.
func NewTransactionmodelFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionmodelFilterer, error) {
	contract, err := bindTransactionmodel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionmodelFilterer{contract: contract}, nil
}

// bindTransactionmodel binds a generic wrapper to an already deployed contract.
func bindTransactionmodel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransactionmodelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionmodel *TransactionmodelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transactionmodel.Contract.TransactionmodelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionmodel *TransactionmodelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactionmodel.Contract.TransactionmodelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionmodel *TransactionmodelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactionmodel.Contract.TransactionmodelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactionmodel *TransactionmodelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transactionmodel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactionmodel *TransactionmodelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactionmodel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactionmodel *TransactionmodelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactionmodel.Contract.contract.Transact(opts, method, params...)
}
