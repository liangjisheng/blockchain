// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coachnft

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CoachnftABI is the input ABI used to generate the binding from.
const CoachnftABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"batchAddWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"fans\",\"type\":\"address[]\"}],\"name\":\"batchMintToFans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"batchRemoveWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedFans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fansClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipClaimState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintToFans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"reserveCoachNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Coachnft is an auto generated Go binding around an Ethereum contract.
type Coachnft struct {
	CoachnftCaller     // Read-only binding to the contract
	CoachnftTransactor // Write-only binding to the contract
	CoachnftFilterer   // Log filterer for contract events
}

// CoachnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoachnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoachnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoachnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoachnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoachnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoachnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoachnftSession struct {
	Contract     *Coachnft         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoachnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoachnftCallerSession struct {
	Contract *CoachnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CoachnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoachnftTransactorSession struct {
	Contract     *CoachnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CoachnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoachnftRaw struct {
	Contract *Coachnft // Generic contract binding to access the raw methods on
}

// CoachnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoachnftCallerRaw struct {
	Contract *CoachnftCaller // Generic read-only contract binding to access the raw methods on
}

// CoachnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoachnftTransactorRaw struct {
	Contract *CoachnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoachnft creates a new instance of Coachnft, bound to a specific deployed contract.
func NewCoachnft(address common.Address, backend bind.ContractBackend) (*Coachnft, error) {
	contract, err := bindCoachnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coachnft{CoachnftCaller: CoachnftCaller{contract: contract}, CoachnftTransactor: CoachnftTransactor{contract: contract}, CoachnftFilterer: CoachnftFilterer{contract: contract}}, nil
}

// NewCoachnftCaller creates a new read-only instance of Coachnft, bound to a specific deployed contract.
func NewCoachnftCaller(address common.Address, caller bind.ContractCaller) (*CoachnftCaller, error) {
	contract, err := bindCoachnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoachnftCaller{contract: contract}, nil
}

// NewCoachnftTransactor creates a new write-only instance of Coachnft, bound to a specific deployed contract.
func NewCoachnftTransactor(address common.Address, transactor bind.ContractTransactor) (*CoachnftTransactor, error) {
	contract, err := bindCoachnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoachnftTransactor{contract: contract}, nil
}

// NewCoachnftFilterer creates a new log filterer instance of Coachnft, bound to a specific deployed contract.
func NewCoachnftFilterer(address common.Address, filterer bind.ContractFilterer) (*CoachnftFilterer, error) {
	contract, err := bindCoachnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoachnftFilterer{contract: contract}, nil
}

// bindCoachnft binds a generic wrapper to an already deployed contract.
func bindCoachnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoachnftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coachnft *CoachnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coachnft.Contract.CoachnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coachnft *CoachnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coachnft.Contract.CoachnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coachnft *CoachnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coachnft.Contract.CoachnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coachnft *CoachnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coachnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coachnft *CoachnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coachnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coachnft *CoachnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coachnft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coachnft *CoachnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coachnft *CoachnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coachnft.Contract.BalanceOf(&_Coachnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coachnft *CoachnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coachnft.Contract.BalanceOf(&_Coachnft.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Coachnft *CoachnftCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Coachnft *CoachnftSession) BaseURI() (string, error) {
	return _Coachnft.Contract.BaseURI(&_Coachnft.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Coachnft *CoachnftCallerSession) BaseURI() (string, error) {
	return _Coachnft.Contract.BaseURI(&_Coachnft.CallOpts)
}

// ClaimIsActive is a free data retrieval call binding the contract method 0x5303f68c.
//
// Solidity: function claimIsActive() view returns(bool)
func (_Coachnft *CoachnftCaller) ClaimIsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "claimIsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimIsActive is a free data retrieval call binding the contract method 0x5303f68c.
//
// Solidity: function claimIsActive() view returns(bool)
func (_Coachnft *CoachnftSession) ClaimIsActive() (bool, error) {
	return _Coachnft.Contract.ClaimIsActive(&_Coachnft.CallOpts)
}

// ClaimIsActive is a free data retrieval call binding the contract method 0x5303f68c.
//
// Solidity: function claimIsActive() view returns(bool)
func (_Coachnft *CoachnftCallerSession) ClaimIsActive() (bool, error) {
	return _Coachnft.Contract.ClaimIsActive(&_Coachnft.CallOpts)
}

// ClaimedFans is a free data retrieval call binding the contract method 0x0e56fbe5.
//
// Solidity: function claimedFans(address ) view returns(bool)
func (_Coachnft *CoachnftCaller) ClaimedFans(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "claimedFans", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimedFans is a free data retrieval call binding the contract method 0x0e56fbe5.
//
// Solidity: function claimedFans(address ) view returns(bool)
func (_Coachnft *CoachnftSession) ClaimedFans(arg0 common.Address) (bool, error) {
	return _Coachnft.Contract.ClaimedFans(&_Coachnft.CallOpts, arg0)
}

// ClaimedFans is a free data retrieval call binding the contract method 0x0e56fbe5.
//
// Solidity: function claimedFans(address ) view returns(bool)
func (_Coachnft *CoachnftCallerSession) ClaimedFans(arg0 common.Address) (bool, error) {
	return _Coachnft.Contract.ClaimedFans(&_Coachnft.CallOpts, arg0)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coachnft *CoachnftCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coachnft *CoachnftSession) ContractURI() (string, error) {
	return _Coachnft.Contract.ContractURI(&_Coachnft.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coachnft *CoachnftCallerSession) ContractURI() (string, error) {
	return _Coachnft.Contract.ContractURI(&_Coachnft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Coachnft.Contract.GetApproved(&_Coachnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Coachnft.Contract.GetApproved(&_Coachnft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Coachnft *CoachnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Coachnft *CoachnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Coachnft.Contract.IsApprovedForAll(&_Coachnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Coachnft *CoachnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Coachnft.Contract.IsApprovedForAll(&_Coachnft.CallOpts, owner, operator)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Coachnft *CoachnftCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Coachnft *CoachnftSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Coachnft.Contract.IsWhitelisted(&_Coachnft.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_Coachnft *CoachnftCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _Coachnft.Contract.IsWhitelisted(&_Coachnft.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coachnft *CoachnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coachnft *CoachnftSession) Name() (string, error) {
	return _Coachnft.Contract.Name(&_Coachnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coachnft *CoachnftCallerSession) Name() (string, error) {
	return _Coachnft.Contract.Name(&_Coachnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coachnft *CoachnftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coachnft *CoachnftSession) Owner() (common.Address, error) {
	return _Coachnft.Contract.Owner(&_Coachnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coachnft *CoachnftCallerSession) Owner() (common.Address, error) {
	return _Coachnft.Contract.Owner(&_Coachnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Coachnft.Contract.OwnerOf(&_Coachnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Coachnft *CoachnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Coachnft.Contract.OwnerOf(&_Coachnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coachnft *CoachnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coachnft *CoachnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coachnft.Contract.SupportsInterface(&_Coachnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coachnft *CoachnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coachnft.Contract.SupportsInterface(&_Coachnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coachnft *CoachnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coachnft *CoachnftSession) Symbol() (string, error) {
	return _Coachnft.Contract.Symbol(&_Coachnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coachnft *CoachnftCallerSession) Symbol() (string, error) {
	return _Coachnft.Contract.Symbol(&_Coachnft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Coachnft *CoachnftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Coachnft *CoachnftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Coachnft.Contract.TokenByIndex(&_Coachnft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Coachnft *CoachnftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Coachnft.Contract.TokenByIndex(&_Coachnft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Coachnft *CoachnftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Coachnft *CoachnftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Coachnft.Contract.TokenOfOwnerByIndex(&_Coachnft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Coachnft *CoachnftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Coachnft.Contract.TokenOfOwnerByIndex(&_Coachnft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coachnft *CoachnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coachnft *CoachnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coachnft.Contract.TokenURI(&_Coachnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coachnft *CoachnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coachnft.Contract.TokenURI(&_Coachnft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coachnft *CoachnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coachnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coachnft *CoachnftSession) TotalSupply() (*big.Int, error) {
	return _Coachnft.Contract.TotalSupply(&_Coachnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coachnft *CoachnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Coachnft.Contract.TotalSupply(&_Coachnft.CallOpts)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Coachnft *CoachnftTransactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "addWhitelisted", _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Coachnft *CoachnftSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.AddWhitelisted(&_Coachnft.TransactOpts, _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Coachnft *CoachnftTransactorSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.AddWhitelisted(&_Coachnft.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.Approve(&_Coachnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.Approve(&_Coachnft.TransactOpts, to, tokenId)
}

// BatchAddWhitelisted is a paid mutator transaction binding the contract method 0xd3527bd6.
//
// Solidity: function batchAddWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftTransactor) BatchAddWhitelisted(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "batchAddWhitelisted", _addresses)
}

// BatchAddWhitelisted is a paid mutator transaction binding the contract method 0xd3527bd6.
//
// Solidity: function batchAddWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftSession) BatchAddWhitelisted(_addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchAddWhitelisted(&_Coachnft.TransactOpts, _addresses)
}

// BatchAddWhitelisted is a paid mutator transaction binding the contract method 0xd3527bd6.
//
// Solidity: function batchAddWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftTransactorSession) BatchAddWhitelisted(_addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchAddWhitelisted(&_Coachnft.TransactOpts, _addresses)
}

// BatchMintToFans is a paid mutator transaction binding the contract method 0xf2f594cc.
//
// Solidity: function batchMintToFans(address[] fans) returns()
func (_Coachnft *CoachnftTransactor) BatchMintToFans(opts *bind.TransactOpts, fans []common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "batchMintToFans", fans)
}

// BatchMintToFans is a paid mutator transaction binding the contract method 0xf2f594cc.
//
// Solidity: function batchMintToFans(address[] fans) returns()
func (_Coachnft *CoachnftSession) BatchMintToFans(fans []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchMintToFans(&_Coachnft.TransactOpts, fans)
}

// BatchMintToFans is a paid mutator transaction binding the contract method 0xf2f594cc.
//
// Solidity: function batchMintToFans(address[] fans) returns()
func (_Coachnft *CoachnftTransactorSession) BatchMintToFans(fans []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchMintToFans(&_Coachnft.TransactOpts, fans)
}

// BatchRemoveWhitelisted is a paid mutator transaction binding the contract method 0x766daf7b.
//
// Solidity: function batchRemoveWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftTransactor) BatchRemoveWhitelisted(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "batchRemoveWhitelisted", _addresses)
}

// BatchRemoveWhitelisted is a paid mutator transaction binding the contract method 0x766daf7b.
//
// Solidity: function batchRemoveWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftSession) BatchRemoveWhitelisted(_addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchRemoveWhitelisted(&_Coachnft.TransactOpts, _addresses)
}

// BatchRemoveWhitelisted is a paid mutator transaction binding the contract method 0x766daf7b.
//
// Solidity: function batchRemoveWhitelisted(address[] _addresses) returns()
func (_Coachnft *CoachnftTransactorSession) BatchRemoveWhitelisted(_addresses []common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.BatchRemoveWhitelisted(&_Coachnft.TransactOpts, _addresses)
}

// FansClaim is a paid mutator transaction binding the contract method 0x0d4b5860.
//
// Solidity: function fansClaim() returns()
func (_Coachnft *CoachnftTransactor) FansClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "fansClaim")
}

// FansClaim is a paid mutator transaction binding the contract method 0x0d4b5860.
//
// Solidity: function fansClaim() returns()
func (_Coachnft *CoachnftSession) FansClaim() (*types.Transaction, error) {
	return _Coachnft.Contract.FansClaim(&_Coachnft.TransactOpts)
}

// FansClaim is a paid mutator transaction binding the contract method 0x0d4b5860.
//
// Solidity: function fansClaim() returns()
func (_Coachnft *CoachnftTransactorSession) FansClaim() (*types.Transaction, error) {
	return _Coachnft.Contract.FansClaim(&_Coachnft.TransactOpts)
}

// FlipClaimState is a paid mutator transaction binding the contract method 0x6d60e6c1.
//
// Solidity: function flipClaimState() returns()
func (_Coachnft *CoachnftTransactor) FlipClaimState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "flipClaimState")
}

// FlipClaimState is a paid mutator transaction binding the contract method 0x6d60e6c1.
//
// Solidity: function flipClaimState() returns()
func (_Coachnft *CoachnftSession) FlipClaimState() (*types.Transaction, error) {
	return _Coachnft.Contract.FlipClaimState(&_Coachnft.TransactOpts)
}

// FlipClaimState is a paid mutator transaction binding the contract method 0x6d60e6c1.
//
// Solidity: function flipClaimState() returns()
func (_Coachnft *CoachnftTransactorSession) FlipClaimState() (*types.Transaction, error) {
	return _Coachnft.Contract.FlipClaimState(&_Coachnft.TransactOpts)
}

// MintToFans is a paid mutator transaction binding the contract method 0xfe97059d.
//
// Solidity: function mintToFans(address to) returns()
func (_Coachnft *CoachnftTransactor) MintToFans(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "mintToFans", to)
}

// MintToFans is a paid mutator transaction binding the contract method 0xfe97059d.
//
// Solidity: function mintToFans(address to) returns()
func (_Coachnft *CoachnftSession) MintToFans(to common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.MintToFans(&_Coachnft.TransactOpts, to)
}

// MintToFans is a paid mutator transaction binding the contract method 0xfe97059d.
//
// Solidity: function mintToFans(address to) returns()
func (_Coachnft *CoachnftTransactorSession) MintToFans(to common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.MintToFans(&_Coachnft.TransactOpts, to)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Coachnft *CoachnftTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Coachnft *CoachnftSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.RemoveWhitelisted(&_Coachnft.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Coachnft *CoachnftTransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.RemoveWhitelisted(&_Coachnft.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coachnft *CoachnftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coachnft *CoachnftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coachnft.Contract.RenounceOwnership(&_Coachnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coachnft *CoachnftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coachnft.Contract.RenounceOwnership(&_Coachnft.TransactOpts)
}

// ReserveCoachNFTs is a paid mutator transaction binding the contract method 0x7f7773f2.
//
// Solidity: function reserveCoachNFTs(uint256 numberOfTokens) returns()
func (_Coachnft *CoachnftTransactor) ReserveCoachNFTs(opts *bind.TransactOpts, numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "reserveCoachNFTs", numberOfTokens)
}

// ReserveCoachNFTs is a paid mutator transaction binding the contract method 0x7f7773f2.
//
// Solidity: function reserveCoachNFTs(uint256 numberOfTokens) returns()
func (_Coachnft *CoachnftSession) ReserveCoachNFTs(numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.ReserveCoachNFTs(&_Coachnft.TransactOpts, numberOfTokens)
}

// ReserveCoachNFTs is a paid mutator transaction binding the contract method 0x7f7773f2.
//
// Solidity: function reserveCoachNFTs(uint256 numberOfTokens) returns()
func (_Coachnft *CoachnftTransactorSession) ReserveCoachNFTs(numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.ReserveCoachNFTs(&_Coachnft.TransactOpts, numberOfTokens)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.SafeTransferFrom(&_Coachnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.SafeTransferFrom(&_Coachnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Coachnft *CoachnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Coachnft *CoachnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Coachnft.Contract.SafeTransferFrom0(&_Coachnft.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Coachnft *CoachnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Coachnft.Contract.SafeTransferFrom0(&_Coachnft.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coachnft *CoachnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coachnft *CoachnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coachnft.Contract.SetApprovalForAll(&_Coachnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coachnft *CoachnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coachnft.Contract.SetApprovalForAll(&_Coachnft.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Coachnft *CoachnftTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Coachnft *CoachnftSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Coachnft.Contract.SetBaseURI(&_Coachnft.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Coachnft *CoachnftTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Coachnft.Contract.SetBaseURI(&_Coachnft.TransactOpts, baseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Coachnft *CoachnftTransactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Coachnft *CoachnftSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Coachnft.Contract.SetContractURI(&_Coachnft.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Coachnft *CoachnftTransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Coachnft.Contract.SetContractURI(&_Coachnft.TransactOpts, contractURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.TransferFrom(&_Coachnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Coachnft *CoachnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Coachnft.Contract.TransferFrom(&_Coachnft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coachnft *CoachnftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Coachnft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coachnft *CoachnftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.TransferOwnership(&_Coachnft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coachnft *CoachnftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coachnft.Contract.TransferOwnership(&_Coachnft.TransactOpts, newOwner)
}

// CoachnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Coachnft contract.
type CoachnftApprovalIterator struct {
	Event *CoachnftApproval // Event containing the contract specifics and raw log

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
func (it *CoachnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoachnftApproval)
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
		it.Event = new(CoachnftApproval)
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
func (it *CoachnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoachnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoachnftApproval represents a Approval event raised by the Coachnft contract.
type CoachnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CoachnftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Coachnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CoachnftApprovalIterator{contract: _Coachnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CoachnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Coachnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoachnftApproval)
				if err := _Coachnft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) ParseApproval(log types.Log) (*CoachnftApproval, error) {
	event := new(CoachnftApproval)
	if err := _Coachnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoachnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Coachnft contract.
type CoachnftApprovalForAllIterator struct {
	Event *CoachnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CoachnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoachnftApprovalForAll)
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
		it.Event = new(CoachnftApprovalForAll)
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
func (it *CoachnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoachnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoachnftApprovalForAll represents a ApprovalForAll event raised by the Coachnft contract.
type CoachnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coachnft *CoachnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CoachnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coachnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CoachnftApprovalForAllIterator{contract: _Coachnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coachnft *CoachnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CoachnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coachnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoachnftApprovalForAll)
				if err := _Coachnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coachnft *CoachnftFilterer) ParseApprovalForAll(log types.Log) (*CoachnftApprovalForAll, error) {
	event := new(CoachnftApprovalForAll)
	if err := _Coachnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoachnftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Coachnft contract.
type CoachnftOwnershipTransferredIterator struct {
	Event *CoachnftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CoachnftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoachnftOwnershipTransferred)
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
		it.Event = new(CoachnftOwnershipTransferred)
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
func (it *CoachnftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoachnftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoachnftOwnershipTransferred represents a OwnershipTransferred event raised by the Coachnft contract.
type CoachnftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coachnft *CoachnftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoachnftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coachnft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoachnftOwnershipTransferredIterator{contract: _Coachnft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coachnft *CoachnftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoachnftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coachnft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoachnftOwnershipTransferred)
				if err := _Coachnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Coachnft *CoachnftFilterer) ParseOwnershipTransferred(log types.Log) (*CoachnftOwnershipTransferred, error) {
	event := new(CoachnftOwnershipTransferred)
	if err := _Coachnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CoachnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Coachnft contract.
type CoachnftTransferIterator struct {
	Event *CoachnftTransfer // Event containing the contract specifics and raw log

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
func (it *CoachnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoachnftTransfer)
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
		it.Event = new(CoachnftTransfer)
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
func (it *CoachnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoachnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoachnftTransfer represents a Transfer event raised by the Coachnft contract.
type CoachnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CoachnftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Coachnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CoachnftTransferIterator{contract: _Coachnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CoachnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Coachnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoachnftTransfer)
				if err := _Coachnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Coachnft *CoachnftFilterer) ParseTransfer(log types.Log) (*CoachnftTransfer, error) {
	event := new(CoachnftTransfer)
	if err := _Coachnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
