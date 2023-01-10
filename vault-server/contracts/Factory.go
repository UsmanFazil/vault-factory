// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_USDC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_withdrawAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fundCollector\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawAdminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFundCollector\",\"type\":\"address\"}],\"name\":\"setFundCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWithdrawAdmin\",\"type\":\"address\"}],\"name\":\"setWithdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161128338038061128383398101604081905261002f916100ff565b61003833610093565b600480546001600160a01b03199081166001600160a01b03978816179091556002805482169587169590951790945560038054851692861692909217909155600580549093169316929092179055600155600060065561015d565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146100fa57600080fd5b919050565b600080600080600060a0868803121561011757600080fd5b610120866100e3565b945061012e602087016100e3565b935060408601519250610143606087016100e3565b9150610151608087016100e3565b90509295509295909350565b6111178061016c6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806369fe0e2d116100975780638da5cb5b116100665780638da5cb5b146101fe578063ced72f871461020f578063e74b981b14610217578063f2fde38b1461022a57600080fd5b806369fe0e2d146101a3578063715018a6146101b65780638079e678146101be5780638d654023146101e757600080fd5b80634ccb20c0116100d35780634ccb20c0146101665780635bd05a23146101775780635d12928b1461018857806365486a3c1461019057600080fd5b80633263b545146100fa578063337807861461014057806347b6070814610151575b600080fd5b61012361010836600461058a565b6008602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6003546001600160a01b0316610123565b61016461015f3660046105a3565b61023d565b005b6002546001600160a01b0316610123565b6005546001600160a01b0316610123565b610164610267565b61016461019e3660046105a3565b6103e5565b6101646101b136600461058a565b61040f565b61016461041c565b6101236101cc3660046105a3565b6007602052600090815260409020546001600160a01b031681565b6101f060065481565b604051908152602001610137565b6000546001600160a01b0316610123565b6001546101f0565b6101646102253660046105a3565b610430565b6101646102383660046105a3565b61045a565b6102456104d3565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b336000908152600760205260409020546001600160a01b0316156102de5760405162461bcd60e51b8152602060048201526024808201527f596f75206861766520616c7265616479206465706c6f796564206120636f6e746044820152631c9858dd60e21b60648201526084015b60405180910390fd5b6004546040516000916001600160a01b03169030906102fc9061057d565b6001600160a01b03928316815291166020820152604001604051809103906000f08015801561032f573d6000803e3d6000fd5b5060405163f2fde38b60e01b81523360048201529091506001600160a01b0382169063f2fde38b90602401600060405180830381600087803b15801561037457600080fd5b505af1158015610388573d6000803e3d6000fd5b50505050600654600161039b91906105d3565b600690815533600090815260076020908152604080832080546001600160a01b039096166001600160a01b0319968716811790915593548352600890915290208054909216179055565b6103ed6104d3565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6104176104d3565b600155565b6104246104d3565b61042e600061052d565b565b6104386104d3565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6104626104d3565b6001600160a01b0381166104c75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102d5565b6104d08161052d565b50565b6000546001600160a01b0316331461042e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d5565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610ae7806105fb83390190565b60006020828403121561059c57600080fd5b5035919050565b6000602082840312156105b557600080fd5b81356001600160a01b03811681146105cc57600080fd5b9392505050565b808201808211156105f457634e487b7160e01b600052601160045260246000fd5b9291505056fe608060405234801561001057600080fd5b50604051610ae7380380610ae783398101604081905261002f916100d5565b61003833610069565b600180546001600160a01b039384166001600160a01b03199182161790915560028054929093169116179055610108565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146100d057600080fd5b919050565b600080604083850312156100e857600080fd5b6100f1836100b9565b91506100ff602084016100b9565b90509250929050565b6109d0806101176000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b6b55f251161005b578063b6b55f25146100ea578063c83dd231146100fd578063f18d20be14610110578063f2fde38b1461011857600080fd5b80632e1a7d4d1461008d578063715018a6146100a257806389a30271146100aa5780638da5cb5b146100d9575b600080fd5b6100a061009b366004610888565b61012b565b005b6100a06103f5565b6001546100bd906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b6000546001600160a01b03166100bd565b6100a06100f8366004610888565b610409565b6002546100bd906001600160a01b031681565b6100a0610542565b6100a06101263660046108b6565b610768565b6101336107de565b6001546040516370a0823160e01b815230600482015282916001600160a01b0316906370a0823190602401602060405180830381865afa15801561017b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061019f91906108da565b10156101e95760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b60448201526064015b60405180910390fd5b60006064600260009054906101000a90046001600160a01b03166001600160a01b031663ced72f876040518163ffffffff1660e01b8152600401602060405180830381865afa158015610240573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026491906108da565b61026e9084610909565b6102789190610926565b6001549091506001600160a01b031663a9059cbb336102978486610948565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156102e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610306919061095b565b50600154600254604080516301332c8360e61b815290516001600160a01b039384169363a9059cbb931691634ccb20c09160048083019260209291908290030181865afa15801561035b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061037f919061097d565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303816000875af11580156103cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f0919061095b565b505050565b6103fd6107de565b6104076000610838565b565b6001546040516370a0823160e01b815233600482015282916001600160a01b0316906370a0823190602401602060405180830381865afa158015610451573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047591906108da565b10156104c35760405162461bcd60e51b815260206004820152601960248201527f496e73756666696369656e7420555344432062616c616e63650000000000000060448201526064016101e0565b6001546040516323b872dd60e01b8152336004820152306024820152604481018390526001600160a01b03909116906323b872dd906064016020604051808303816000875af115801561051a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053e919061095b565b5050565b600260009054906101000a90046001600160a01b03166001600160a01b031663337807866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610595573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b9919061097d565b6001600160a01b0316336001600160a01b0316146106125760405162461bcd60e51b8152602060048201526016602482015275496e76616c69642063616c6c6572206164647265737360501b60448201526064016101e0565b60015460025460408051635bd05a2360e01b815290516001600160a01b039384169363a9059cbb931691635bd05a239160048083019260209291908290030181865afa158015610666573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068a919061097d565b6001546040516370a0823160e01b81523060048201526001600160a01b03909116906370a0823190602401602060405180830381865afa1580156106d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f691906108da565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015610741573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610765919061095b565b50565b6107706107de565b6001600160a01b0381166107d55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101e0565b61076581610838565b6000546001600160a01b031633146104075760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561089a57600080fd5b5035919050565b6001600160a01b038116811461076557600080fd5b6000602082840312156108c857600080fd5b81356108d3816108a1565b9392505050565b6000602082840312156108ec57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610920576109206108f3565b92915050565b60008261094357634e487b7160e01b600052601260045260246000fd5b500490565b81810381811115610920576109206108f3565b60006020828403121561096d57600080fd5b815180151581146108d357600080fd5b60006020828403121561098f57600080fd5b81516108d3816108a156fea26469706673582212203e2c90713b7ff126b2c6ec9a899797a9400f9db437be9d0ba12c319b08102cd264736f6c63430008110033a2646970667358221220246a664400b1183cb37e629a7bedb71149023a3ae850e1731c94108990e08c3864736f6c63430008110033",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FactoryMetaData.Bin instead.
var FactoryBin = FactoryMetaData.Bin

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _USDC common.Address, _feeRecipient common.Address, _fee *big.Int, _withdrawAdmin common.Address, _fundCollector common.Address) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryBin), backend, _USDC, _feeRecipient, _fee, _withdrawAdmin, _fundCollector)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_Factory *FactoryCaller) ContractAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "contractAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_Factory *FactorySession) ContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.ContractAddresses(&_Factory.CallOpts, arg0)
}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) ContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.ContractAddresses(&_Factory.CallOpts, arg0)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Factory *FactoryCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Factory *FactorySession) GetFee() (*big.Int, error) {
	return _Factory.Contract.GetFee(&_Factory.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Factory *FactoryCallerSession) GetFee() (*big.Int, error) {
	return _Factory.Contract.GetFee(&_Factory.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_Factory *FactoryCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_Factory *FactorySession) GetFeeRecipient() (common.Address, error) {
	return _Factory.Contract.GetFeeRecipient(&_Factory.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_Factory *FactoryCallerSession) GetFeeRecipient() (common.Address, error) {
	return _Factory.Contract.GetFeeRecipient(&_Factory.CallOpts)
}

// GetFundCollector is a free data retrieval call binding the contract method 0x5bd05a23.
//
// Solidity: function getFundCollector() view returns(address)
func (_Factory *FactoryCaller) GetFundCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getFundCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundCollector is a free data retrieval call binding the contract method 0x5bd05a23.
//
// Solidity: function getFundCollector() view returns(address)
func (_Factory *FactorySession) GetFundCollector() (common.Address, error) {
	return _Factory.Contract.GetFundCollector(&_Factory.CallOpts)
}

// GetFundCollector is a free data retrieval call binding the contract method 0x5bd05a23.
//
// Solidity: function getFundCollector() view returns(address)
func (_Factory *FactoryCallerSession) GetFundCollector() (common.Address, error) {
	return _Factory.Contract.GetFundCollector(&_Factory.CallOpts)
}

// GetWithdrawAdminAddr is a free data retrieval call binding the contract method 0x33780786.
//
// Solidity: function getWithdrawAdminAddr() view returns(address)
func (_Factory *FactoryCaller) GetWithdrawAdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getWithdrawAdminAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawAdminAddr is a free data retrieval call binding the contract method 0x33780786.
//
// Solidity: function getWithdrawAdminAddr() view returns(address)
func (_Factory *FactorySession) GetWithdrawAdminAddr() (common.Address, error) {
	return _Factory.Contract.GetWithdrawAdminAddr(&_Factory.CallOpts)
}

// GetWithdrawAdminAddr is a free data retrieval call binding the contract method 0x33780786.
//
// Solidity: function getWithdrawAdminAddr() view returns(address)
func (_Factory *FactoryCallerSession) GetWithdrawAdminAddr() (common.Address, error) {
	return _Factory.Contract.GetWithdrawAdminAddr(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactorySession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCallerSession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_Factory *FactoryCaller) TotalVaults(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "totalVaults")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_Factory *FactorySession) TotalVaults() (*big.Int, error) {
	return _Factory.Contract.TotalVaults(&_Factory.CallOpts)
}

// TotalVaults is a free data retrieval call binding the contract method 0x8d654023.
//
// Solidity: function totalVaults() view returns(uint256)
func (_Factory *FactoryCallerSession) TotalVaults() (*big.Int, error) {
	return _Factory.Contract.TotalVaults(&_Factory.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x8079e678.
//
// Solidity: function userContract(address ) view returns(address)
func (_Factory *FactoryCaller) UserContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "userContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContract is a free data retrieval call binding the contract method 0x8079e678.
//
// Solidity: function userContract(address ) view returns(address)
func (_Factory *FactorySession) UserContract(arg0 common.Address) (common.Address, error) {
	return _Factory.Contract.UserContract(&_Factory.CallOpts, arg0)
}

// UserContract is a free data retrieval call binding the contract method 0x8079e678.
//
// Solidity: function userContract(address ) view returns(address)
func (_Factory *FactoryCallerSession) UserContract(arg0 common.Address) (common.Address, error) {
	return _Factory.Contract.UserContract(&_Factory.CallOpts, arg0)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns()
func (_Factory *FactoryTransactor) CreateVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createVault")
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns()
func (_Factory *FactorySession) CreateVault() (*types.Transaction, error) {
	return _Factory.Contract.CreateVault(&_Factory.TransactOpts)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns()
func (_Factory *FactoryTransactorSession) CreateVault() (*types.Transaction, error) {
	return _Factory.Contract.CreateVault(&_Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Factory *FactoryTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Factory *FactorySession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.SetFee(&_Factory.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Factory *FactoryTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.SetFee(&_Factory.TransactOpts, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Factory *FactoryTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Factory *FactorySession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeRecipient(&_Factory.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Factory *FactoryTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeRecipient(&_Factory.TransactOpts, newFeeRecipient)
}

// SetFundCollector is a paid mutator transaction binding the contract method 0x47b60708.
//
// Solidity: function setFundCollector(address newFundCollector) returns()
func (_Factory *FactoryTransactor) SetFundCollector(opts *bind.TransactOpts, newFundCollector common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFundCollector", newFundCollector)
}

// SetFundCollector is a paid mutator transaction binding the contract method 0x47b60708.
//
// Solidity: function setFundCollector(address newFundCollector) returns()
func (_Factory *FactorySession) SetFundCollector(newFundCollector common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFundCollector(&_Factory.TransactOpts, newFundCollector)
}

// SetFundCollector is a paid mutator transaction binding the contract method 0x47b60708.
//
// Solidity: function setFundCollector(address newFundCollector) returns()
func (_Factory *FactoryTransactorSession) SetFundCollector(newFundCollector common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFundCollector(&_Factory.TransactOpts, newFundCollector)
}

// SetWithdrawAdmin is a paid mutator transaction binding the contract method 0x65486a3c.
//
// Solidity: function setWithdrawAdmin(address newWithdrawAdmin) returns()
func (_Factory *FactoryTransactor) SetWithdrawAdmin(opts *bind.TransactOpts, newWithdrawAdmin common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setWithdrawAdmin", newWithdrawAdmin)
}

// SetWithdrawAdmin is a paid mutator transaction binding the contract method 0x65486a3c.
//
// Solidity: function setWithdrawAdmin(address newWithdrawAdmin) returns()
func (_Factory *FactorySession) SetWithdrawAdmin(newWithdrawAdmin common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetWithdrawAdmin(&_Factory.TransactOpts, newWithdrawAdmin)
}

// SetWithdrawAdmin is a paid mutator transaction binding the contract method 0x65486a3c.
//
// Solidity: function setWithdrawAdmin(address newWithdrawAdmin) returns()
func (_Factory *FactoryTransactorSession) SetWithdrawAdmin(newWithdrawAdmin common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetWithdrawAdmin(&_Factory.TransactOpts, newWithdrawAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Factory contract.
type FactoryOwnershipTransferredIterator struct {
	Event *FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryOwnershipTransferred)
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
		it.Event = new(FactoryOwnershipTransferred)
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
func (it *FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Factory contract.
type FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryOwnershipTransferredIterator{contract: _Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryOwnershipTransferred)
				if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Factory *FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*FactoryOwnershipTransferred, error) {
	event := new(FactoryOwnershipTransferred)
	if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
