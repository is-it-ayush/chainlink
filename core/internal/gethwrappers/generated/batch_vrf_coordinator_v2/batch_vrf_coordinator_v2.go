// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batch_vrf_coordinator_v2

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

type VRFTypesProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type VRFTypesRequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
}

var BatchVRFCoordinatorV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"COORDINATOR\",\"outputs\":[{\"internalType\":\"contractVRFCoordinatorV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRFTypes.Proof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structVRFTypes.RequestCommitment[]\",\"name\":\"rcs\",\"type\":\"tuple[]\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161088838038061088883398101604081905261002f91610044565b60601b6001600160601b031916608052610074565b60006020828403121561005657600080fd5b81516001600160a01b038116811461006d57600080fd5b9392505050565b60805160601c6107f0610098600039600081816055015261011d01526107f06000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806308b2da0a1461003b5780633b2bcbf114610050575b600080fd5b61004e6100493660046103b4565b6100a0565b005b6100777f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b805182511461010f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f696e70757420617272617920617267206c656e67746873206d69736d61746368604482015260640160405180910390fd5b60005b825181101561020d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663af198b9784838151811061016957610169610785565b602002602001015184848151811061018357610183610785565b60200260200101516040518363ffffffff1660e01b81526004016101a8929190610568565b602060405180830381600087803b1580156101c257600080fd5b505af11580156101d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101fa919061050a565b508061020581610725565b915050610112565b505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461023657600080fd5b919050565b600082601f83011261024c57600080fd5b8135602061026161025c83610701565b6106b2565b8281528181019085830160a08086028801850189101561028057600080fd5b6000805b878110156103015782848c03121561029a578182fd5b6102a2610665565b6102ab8561039c565b81526102b888860161039c565b8882015260406102c9818701610388565b9082015260606102da868201610388565b9082015260806102eb868201610212565b9082015286529486019492820192600101610284565b50929998505050505050505050565b600082601f83011261032157600080fd5b6040516040810181811067ffffffffffffffff82111715610344576103446107b4565b806040525080838560408601111561035b57600080fd5b60005b600281101561037d57813583526020928301929091019060010161035e565b509195945050505050565b803563ffffffff8116811461023657600080fd5b803567ffffffffffffffff8116811461023657600080fd5b60008060408084860312156103c857600080fd5b833567ffffffffffffffff808211156103e057600080fd5b818601915086601f8301126103f457600080fd5b8135602061040461025c83610701565b828152818101908583016101a0808602880185018d101561042457600080fd5b600097505b858810156104d95780828e03121561044057600080fd5b61044861068e565b6104528e84610310565b81526104608e8b8501610310565b8682015260808301358a82015260a0830135606082015260c0830135608082015261048d60e08401610212565b60a08201526101006104a18f828601610310565b60c08301526104b48f6101408601610310565b60e0830152610180840135908201528452600197909701969284019290810190610429565b509098505050870135935050808311156104f257600080fd5b50506105008582860161023b565b9150509250929050565b60006020828403121561051c57600080fd5b81516bffffffffffffffffffffffff8116811461053857600080fd5b9392505050565b8060005b6002811015610562578151845260209384019390910190600101610543565b50505050565b60006102408201905061057c82855161053f565b602084015161058e604084018261053f565b5060408401516080830152606084015160a0830152608084015160c083015273ffffffffffffffffffffffffffffffffffffffff60a08501511660e083015260c08401516101006105e18185018361053f565b60e086015191506105f661014085018361053f565b85015161018084015250825167ffffffffffffffff9081166101a08401526020840151166101c0830152604083015163ffffffff9081166101e0840152606084015116610200830152608083015173ffffffffffffffffffffffffffffffffffffffff16610220830152610538565b60405160a0810167ffffffffffffffff81118282101715610688576106886107b4565b60405290565b604051610120810167ffffffffffffffff81118282101715610688576106886107b4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156106f9576106f96107b4565b604052919050565b600067ffffffffffffffff82111561071b5761071b6107b4565b5060051b60200190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561077e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var BatchVRFCoordinatorV2ABI = BatchVRFCoordinatorV2MetaData.ABI

var BatchVRFCoordinatorV2Bin = BatchVRFCoordinatorV2MetaData.Bin

func DeployBatchVRFCoordinatorV2(auth *bind.TransactOpts, backend bind.ContractBackend, coordinatorAddr common.Address) (common.Address, *types.Transaction, *BatchVRFCoordinatorV2, error) {
	parsed, err := BatchVRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchVRFCoordinatorV2Bin), backend, coordinatorAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchVRFCoordinatorV2{BatchVRFCoordinatorV2Caller: BatchVRFCoordinatorV2Caller{contract: contract}, BatchVRFCoordinatorV2Transactor: BatchVRFCoordinatorV2Transactor{contract: contract}, BatchVRFCoordinatorV2Filterer: BatchVRFCoordinatorV2Filterer{contract: contract}}, nil
}

type BatchVRFCoordinatorV2 struct {
	address common.Address
	abi     abi.ABI
	BatchVRFCoordinatorV2Caller
	BatchVRFCoordinatorV2Transactor
	BatchVRFCoordinatorV2Filterer
}

type BatchVRFCoordinatorV2Caller struct {
	contract *bind.BoundContract
}

type BatchVRFCoordinatorV2Transactor struct {
	contract *bind.BoundContract
}

type BatchVRFCoordinatorV2Filterer struct {
	contract *bind.BoundContract
}

type BatchVRFCoordinatorV2Session struct {
	Contract     *BatchVRFCoordinatorV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BatchVRFCoordinatorV2CallerSession struct {
	Contract *BatchVRFCoordinatorV2Caller
	CallOpts bind.CallOpts
}

type BatchVRFCoordinatorV2TransactorSession struct {
	Contract     *BatchVRFCoordinatorV2Transactor
	TransactOpts bind.TransactOpts
}

type BatchVRFCoordinatorV2Raw struct {
	Contract *BatchVRFCoordinatorV2
}

type BatchVRFCoordinatorV2CallerRaw struct {
	Contract *BatchVRFCoordinatorV2Caller
}

type BatchVRFCoordinatorV2TransactorRaw struct {
	Contract *BatchVRFCoordinatorV2Transactor
}

func NewBatchVRFCoordinatorV2(address common.Address, backend bind.ContractBackend) (*BatchVRFCoordinatorV2, error) {
	abi, err := abi.JSON(strings.NewReader(BatchVRFCoordinatorV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBatchVRFCoordinatorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchVRFCoordinatorV2{address: address, abi: abi, BatchVRFCoordinatorV2Caller: BatchVRFCoordinatorV2Caller{contract: contract}, BatchVRFCoordinatorV2Transactor: BatchVRFCoordinatorV2Transactor{contract: contract}, BatchVRFCoordinatorV2Filterer: BatchVRFCoordinatorV2Filterer{contract: contract}}, nil
}

func NewBatchVRFCoordinatorV2Caller(address common.Address, caller bind.ContractCaller) (*BatchVRFCoordinatorV2Caller, error) {
	contract, err := bindBatchVRFCoordinatorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVRFCoordinatorV2Caller{contract: contract}, nil
}

func NewBatchVRFCoordinatorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BatchVRFCoordinatorV2Transactor, error) {
	contract, err := bindBatchVRFCoordinatorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVRFCoordinatorV2Transactor{contract: contract}, nil
}

func NewBatchVRFCoordinatorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BatchVRFCoordinatorV2Filterer, error) {
	contract, err := bindBatchVRFCoordinatorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchVRFCoordinatorV2Filterer{contract: contract}, nil
}

func bindBatchVRFCoordinatorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchVRFCoordinatorV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVRFCoordinatorV2.Contract.BatchVRFCoordinatorV2Caller.contract.Call(opts, result, method, params...)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.BatchVRFCoordinatorV2Transactor.contract.Transfer(opts)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.BatchVRFCoordinatorV2Transactor.contract.Transact(opts, method, params...)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVRFCoordinatorV2.Contract.contract.Call(opts, result, method, params...)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.contract.Transfer(opts)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.contract.Transact(opts, method, params...)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Caller) COORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchVRFCoordinatorV2.contract.Call(opts, &out, "COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Session) COORDINATOR() (common.Address, error) {
	return _BatchVRFCoordinatorV2.Contract.COORDINATOR(&_BatchVRFCoordinatorV2.CallOpts)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2CallerSession) COORDINATOR() (common.Address, error) {
	return _BatchVRFCoordinatorV2.Contract.COORDINATOR(&_BatchVRFCoordinatorV2.CallOpts)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Transactor) FulfillRandomWords(opts *bind.TransactOpts, proofs []VRFTypesProof, rcs []VRFTypesRequestCommitment) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.contract.Transact(opts, "fulfillRandomWords", proofs, rcs)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2Session) FulfillRandomWords(proofs []VRFTypesProof, rcs []VRFTypesRequestCommitment) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.FulfillRandomWords(&_BatchVRFCoordinatorV2.TransactOpts, proofs, rcs)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2TransactorSession) FulfillRandomWords(proofs []VRFTypesProof, rcs []VRFTypesRequestCommitment) (*types.Transaction, error) {
	return _BatchVRFCoordinatorV2.Contract.FulfillRandomWords(&_BatchVRFCoordinatorV2.TransactOpts, proofs, rcs)
}

func (_BatchVRFCoordinatorV2 *BatchVRFCoordinatorV2) Address() common.Address {
	return _BatchVRFCoordinatorV2.address
}

type BatchVRFCoordinatorV2Interface interface {
	COORDINATOR(opts *bind.CallOpts) (common.Address, error)

	FulfillRandomWords(opts *bind.TransactOpts, proofs []VRFTypesProof, rcs []VRFTypesRequestCommitment) (*types.Transaction, error)

	Address() common.Address
}
