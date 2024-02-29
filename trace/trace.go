// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

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

// UserPassStoreUserPassInfo is an auto generated low-level Go binding around an user-defined struct.
type UserPassStoreUserPassInfo struct {
	Username     string
	Passwordhash string
	Name         string
	Phone        string
	Description  string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwordhash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addorupdateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwordhash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"tracedataByName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwordhash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structUserPassStore.UserPassInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610f9f8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063216eb219146100435780635c7e94951461005f578063e0c1f07514610093575b5f80fd5b61005d60048036038101906100589190610918565b6100c3565b005b61007960048036038101906100749190610a1b565b6101af565b60405161008a959493929190610adc565b60405180910390f35b6100ad60048036038101906100a89190610a1b565b610496565b6040516100ba9190610c27565b60405180910390f35b845f866040516100d39190610c81565b90815260200160405180910390205f0190816100ef9190610e9a565b50835f866040516101009190610c81565b9081526020016040518091039020600101908161011d9190610e9a565b50825f8660405161012e9190610c81565b9081526020016040518091039020600201908161014b9190610e9a565b50815f8660405161015c9190610c81565b908152602001604051809103902060030190816101799190610e9a565b50805f8660405161018a9190610c81565b908152602001604051809103902060040190816101a79190610e9a565b505050505050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546101e590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461021190610cc4565b801561025c5780601f106102335761010080835404028352916020019161025c565b820191905f5260205f20905b81548152906001019060200180831161023f57829003601f168201915b50505050509080600101805461027190610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461029d90610cc4565b80156102e85780601f106102bf576101008083540402835291602001916102e8565b820191905f5260205f20905b8154815290600101906020018083116102cb57829003601f168201915b5050505050908060020180546102fd90610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461032990610cc4565b80156103745780601f1061034b57610100808354040283529160200191610374565b820191905f5260205f20905b81548152906001019060200180831161035757829003601f168201915b50505050509080600301805461038990610cc4565b80601f01602080910402602001604051908101604052809291908181526020018280546103b590610cc4565b80156104005780601f106103d757610100808354040283529160200191610400565b820191905f5260205f20905b8154815290600101906020018083116103e357829003601f168201915b50505050509080600401805461041590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461044190610cc4565b801561048c5780601f106104635761010080835404028352916020019161048c565b820191905f5260205f20905b81548152906001019060200180831161046f57829003601f168201915b5050505050905085565b61049e61079c565b5f826040516104ad9190610c81565b90815260200160405180910390206040518060a00160405290815f820180546104d590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461050190610cc4565b801561054c5780601f106105235761010080835404028352916020019161054c565b820191905f5260205f20905b81548152906001019060200180831161052f57829003601f168201915b5050505050815260200160018201805461056590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461059190610cc4565b80156105dc5780601f106105b3576101008083540402835291602001916105dc565b820191905f5260205f20905b8154815290600101906020018083116105bf57829003601f168201915b505050505081526020016002820180546105f590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461062190610cc4565b801561066c5780601f106106435761010080835404028352916020019161066c565b820191905f5260205f20905b81548152906001019060200180831161064f57829003601f168201915b5050505050815260200160038201805461068590610cc4565b80601f01602080910402602001604051908101604052809291908181526020018280546106b190610cc4565b80156106fc5780601f106106d3576101008083540402835291602001916106fc565b820191905f5260205f20905b8154815290600101906020018083116106df57829003601f168201915b5050505050815260200160048201805461071590610cc4565b80601f016020809104026020016040519081016040528092919081815260200182805461074190610cc4565b801561078c5780601f106107635761010080835404028352916020019161078c565b820191905f5260205f20905b81548152906001019060200180831161076f57829003601f168201915b5050505050815250509050919050565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61082a826107e4565b810181811067ffffffffffffffff82111715610849576108486107f4565b5b80604052505050565b5f61085b6107cb565b90506108678282610821565b919050565b5f67ffffffffffffffff821115610886576108856107f4565b5b61088f826107e4565b9050602081019050919050565b828183375f83830152505050565b5f6108bc6108b78461086c565b610852565b9050828152602081018484840111156108d8576108d76107e0565b5b6108e384828561089c565b509392505050565b5f82601f8301126108ff576108fe6107dc565b5b813561090f8482602086016108aa565b91505092915050565b5f805f805f60a08688031215610931576109306107d4565b5b5f86013567ffffffffffffffff81111561094e5761094d6107d8565b5b61095a888289016108eb565b955050602086013567ffffffffffffffff81111561097b5761097a6107d8565b5b610987888289016108eb565b945050604086013567ffffffffffffffff8111156109a8576109a76107d8565b5b6109b4888289016108eb565b935050606086013567ffffffffffffffff8111156109d5576109d46107d8565b5b6109e1888289016108eb565b925050608086013567ffffffffffffffff811115610a0257610a016107d8565b5b610a0e888289016108eb565b9150509295509295909350565b5f60208284031215610a3057610a2f6107d4565b5b5f82013567ffffffffffffffff811115610a4d57610a4c6107d8565b5b610a59848285016108eb565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610a99578082015181840152602081019050610a7e565b5f8484015250505050565b5f610aae82610a62565b610ab88185610a6c565b9350610ac8818560208601610a7c565b610ad1816107e4565b840191505092915050565b5f60a0820190508181035f830152610af48188610aa4565b90508181036020830152610b088187610aa4565b90508181036040830152610b1c8186610aa4565b90508181036060830152610b308185610aa4565b90508181036080830152610b448184610aa4565b90509695505050505050565b5f82825260208201905092915050565b5f610b6a82610a62565b610b748185610b50565b9350610b84818560208601610a7c565b610b8d816107e4565b840191505092915050565b5f60a083015f8301518482035f860152610bb28282610b60565b91505060208301518482036020860152610bcc8282610b60565b91505060408301518482036040860152610be68282610b60565b91505060608301518482036060860152610c008282610b60565b91505060808301518482036080860152610c1a8282610b60565b9150508091505092915050565b5f6020820190508181035f830152610c3f8184610b98565b905092915050565b5f81905092915050565b5f610c5b82610a62565b610c658185610c47565b9350610c75818560208601610a7c565b80840191505092915050565b5f610c8c8284610c51565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610cdb57607f821691505b602082108103610cee57610ced610c97565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610d507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d15565b610d5a8683610d15565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610d9e610d99610d9484610d72565b610d7b565b610d72565b9050919050565b5f819050919050565b610db783610d84565b610dcb610dc382610da5565b848454610d21565b825550505050565b5f90565b610ddf610dd3565b610dea818484610dae565b505050565b5b81811015610e0d57610e025f82610dd7565b600181019050610df0565b5050565b601f821115610e5257610e2381610cf4565b610e2c84610d06565b81016020851015610e3b578190505b610e4f610e4785610d06565b830182610def565b50505b505050565b5f82821c905092915050565b5f610e725f1984600802610e57565b1980831691505092915050565b5f610e8a8383610e63565b9150826002028217905092915050565b610ea382610a62565b67ffffffffffffffff811115610ebc57610ebb6107f4565b5b610ec68254610cc4565b610ed1828285610e11565b5f60209050601f831160018114610f02575f8415610ef0578287015190505b610efa8582610e7f565b865550610f61565b601f198416610f1086610cf4565b5f5b82811015610f3757848901518255600182019150602085019450602081019050610f12565b86831015610f545784890151610f50601f891682610e63565b8355505b6001600288020188555050505b50505050505056fea26469706673582212202354aa641bfacd7b7d39021f1b899c70f6f947ac4d92f738436c6d00a4d8298d64736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string username, string passwordhash, string name, string phone, string description)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Username     string
	Passwordhash string
	Name         string
	Phone        string
	Description  string
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Username     string
		Passwordhash string
		Name         string
		Phone        string
		Description  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Username = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Passwordhash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Phone = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string username, string passwordhash, string name, string phone, string description)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Username     string
	Passwordhash string
	Name         string
	Phone        string
	Description  string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string username, string passwordhash, string name, string phone, string description)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Username     string
	Passwordhash string
	Name         string
	Phone        string
	Description  string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// TracedataByName is a free data retrieval call binding the contract method 0xe0c1f075.
//
// Solidity: function tracedataByName(string name) view returns((string,string,string,string,string))
func (_Trace *TraceCaller) TracedataByName(opts *bind.CallOpts, name string) (UserPassStoreUserPassInfo, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedataByName", name)

	if err != nil {
		return *new(UserPassStoreUserPassInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UserPassStoreUserPassInfo)).(*UserPassStoreUserPassInfo)

	return out0, err

}

// TracedataByName is a free data retrieval call binding the contract method 0xe0c1f075.
//
// Solidity: function tracedataByName(string name) view returns((string,string,string,string,string))
func (_Trace *TraceSession) TracedataByName(name string) (UserPassStoreUserPassInfo, error) {
	return _Trace.Contract.TracedataByName(&_Trace.CallOpts, name)
}

// TracedataByName is a free data retrieval call binding the contract method 0xe0c1f075.
//
// Solidity: function tracedataByName(string name) view returns((string,string,string,string,string))
func (_Trace *TraceCallerSession) TracedataByName(name string) (UserPassStoreUserPassInfo, error) {
	return _Trace.Contract.TracedataByName(&_Trace.CallOpts, name)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x216eb219.
//
// Solidity: function addorupdateData(string username, string passwordhash, string name, string phone, string description) returns()
func (_Trace *TraceTransactor) AddorupdateData(opts *bind.TransactOpts, username string, passwordhash string, name string, phone string, description string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addorupdateData", username, passwordhash, name, phone, description)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x216eb219.
//
// Solidity: function addorupdateData(string username, string passwordhash, string name, string phone, string description) returns()
func (_Trace *TraceSession) AddorupdateData(username string, passwordhash string, name string, phone string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, username, passwordhash, name, phone, description)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x216eb219.
//
// Solidity: function addorupdateData(string username, string passwordhash, string name, string phone, string description) returns()
func (_Trace *TraceTransactorSession) AddorupdateData(username string, passwordhash string, name string, phone string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, username, passwordhash, name, phone, description)
}
