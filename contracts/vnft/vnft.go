// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vnft

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

// IERC721AUpgradeableTokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721AUpgradeableTokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// VnftMetaData contains all meta data concerning the Vnft contract.
var VnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_before\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_after\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_before\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_after\",\"type\":\"address\"}],\"name\":\"LiquidStakingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeNftsOfStakingPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeValidatorsOfStakingPool\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership\",\"name\":\"ownership\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getActiveNftCountsOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmptyNftCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getEmptyNftCountsOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getNextValidatorWithdrawalCredential\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getNftExitBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveNftCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getUserActiveNftCountsOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getUserNftGasHeight\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getUserNftWithdrawalCredentialOfTokenId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"initHeightOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"lastOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"numberMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorEmptyNftIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorEmptyNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"operatorOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidStakingContractAddress\",\"type\":\"address\"}],\"name\":\"setLiquidStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_exitBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"setNftExitBlockNumbers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setUserNftGasHeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"tokenOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExitButNoBurnNftCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"validatorExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"validatorOf\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"validatorsOfOperator\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"validatorsOfOwner\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"whiteListBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"whiteListMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VnftABI is the input ABI used to generate the binding from.
// Deprecated: Use VnftMetaData.ABI instead.
var VnftABI = VnftMetaData.ABI

// Vnft is an auto generated Go binding around an Ethereum contract.
type Vnft struct {
	VnftCaller     // Read-only binding to the contract
	VnftTransactor // Write-only binding to the contract
	VnftFilterer   // Log filterer for contract events
}

// VnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type VnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VnftSession struct {
	Contract     *Vnft             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VnftCallerSession struct {
	Contract *VnftCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VnftTransactorSession struct {
	Contract     *VnftTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type VnftRaw struct {
	Contract *Vnft // Generic contract binding to access the raw methods on
}

// VnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VnftCallerRaw struct {
	Contract *VnftCaller // Generic read-only contract binding to access the raw methods on
}

// VnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VnftTransactorRaw struct {
	Contract *VnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVnft creates a new instance of Vnft, bound to a specific deployed contract.
func NewVnft(address common.Address, backend bind.ContractBackend) (*Vnft, error) {
	contract, err := bindVnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vnft{VnftCaller: VnftCaller{contract: contract}, VnftTransactor: VnftTransactor{contract: contract}, VnftFilterer: VnftFilterer{contract: contract}}, nil
}

// NewVnftCaller creates a new read-only instance of Vnft, bound to a specific deployed contract.
func NewVnftCaller(address common.Address, caller bind.ContractCaller) (*VnftCaller, error) {
	contract, err := bindVnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VnftCaller{contract: contract}, nil
}

// NewVnftTransactor creates a new write-only instance of Vnft, bound to a specific deployed contract.
func NewVnftTransactor(address common.Address, transactor bind.ContractTransactor) (*VnftTransactor, error) {
	contract, err := bindVnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VnftTransactor{contract: contract}, nil
}

// NewVnftFilterer creates a new log filterer instance of Vnft, bound to a specific deployed contract.
func NewVnftFilterer(address common.Address, filterer bind.ContractFilterer) (*VnftFilterer, error) {
	contract, err := bindVnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VnftFilterer{contract: contract}, nil
}

// bindVnft binds a generic wrapper to an already deployed contract.
func bindVnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VnftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vnft *VnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vnft.Contract.VnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vnft *VnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vnft.Contract.VnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vnft *VnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vnft.Contract.VnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vnft *VnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vnft *VnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vnft *VnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vnft.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Vnft *VnftCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Vnft *VnftSession) MAXSUPPLY() (*big.Int, error) {
	return _Vnft.Contract.MAXSUPPLY(&_Vnft.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Vnft *VnftCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Vnft.Contract.MAXSUPPLY(&_Vnft.CallOpts)
}

// ActiveNftsOfStakingPool is a free data retrieval call binding the contract method 0xe30391e1.
//
// Solidity: function activeNftsOfStakingPool() view returns(uint256[])
func (_Vnft *VnftCaller) ActiveNftsOfStakingPool(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "activeNftsOfStakingPool")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ActiveNftsOfStakingPool is a free data retrieval call binding the contract method 0xe30391e1.
//
// Solidity: function activeNftsOfStakingPool() view returns(uint256[])
func (_Vnft *VnftSession) ActiveNftsOfStakingPool() ([]*big.Int, error) {
	return _Vnft.Contract.ActiveNftsOfStakingPool(&_Vnft.CallOpts)
}

// ActiveNftsOfStakingPool is a free data retrieval call binding the contract method 0xe30391e1.
//
// Solidity: function activeNftsOfStakingPool() view returns(uint256[])
func (_Vnft *VnftCallerSession) ActiveNftsOfStakingPool() ([]*big.Int, error) {
	return _Vnft.Contract.ActiveNftsOfStakingPool(&_Vnft.CallOpts)
}

// ActiveValidatorsOfStakingPool is a free data retrieval call binding the contract method 0xa109fdcf.
//
// Solidity: function activeValidatorsOfStakingPool() view returns(bytes[])
func (_Vnft *VnftCaller) ActiveValidatorsOfStakingPool(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "activeValidatorsOfStakingPool")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ActiveValidatorsOfStakingPool is a free data retrieval call binding the contract method 0xa109fdcf.
//
// Solidity: function activeValidatorsOfStakingPool() view returns(bytes[])
func (_Vnft *VnftSession) ActiveValidatorsOfStakingPool() ([][]byte, error) {
	return _Vnft.Contract.ActiveValidatorsOfStakingPool(&_Vnft.CallOpts)
}

// ActiveValidatorsOfStakingPool is a free data retrieval call binding the contract method 0xa109fdcf.
//
// Solidity: function activeValidatorsOfStakingPool() view returns(bytes[])
func (_Vnft *VnftCallerSession) ActiveValidatorsOfStakingPool() ([][]byte, error) {
	return _Vnft.Contract.ActiveValidatorsOfStakingPool(&_Vnft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vnft *VnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vnft *VnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vnft.Contract.BalanceOf(&_Vnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vnft *VnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vnft.Contract.BalanceOf(&_Vnft.CallOpts, owner)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Vnft *VnftCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721AUpgradeableTokenOwnership)).(*IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Vnft *VnftSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Vnft.Contract.ExplicitOwnershipOf(&_Vnft.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Vnft *VnftCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Vnft.Contract.ExplicitOwnershipOf(&_Vnft.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Vnft *VnftCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721AUpgradeableTokenOwnership)).(*[]IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Vnft *VnftSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Vnft.Contract.ExplicitOwnershipsOf(&_Vnft.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Vnft *VnftCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Vnft.Contract.ExplicitOwnershipsOf(&_Vnft.CallOpts, tokenIds)
}

// GetActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x7094c2a2.
//
// Solidity: function getActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCaller) GetActiveNftCountsOfOperator(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getActiveNftCountsOfOperator", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x7094c2a2.
//
// Solidity: function getActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftSession) GetActiveNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetActiveNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x7094c2a2.
//
// Solidity: function getActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCallerSession) GetActiveNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetActiveNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Vnft *VnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Vnft *VnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.GetApproved(&_Vnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Vnft *VnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.GetApproved(&_Vnft.CallOpts, tokenId)
}

// GetEmptyNftCounts is a free data retrieval call binding the contract method 0x0ec3be08.
//
// Solidity: function getEmptyNftCounts() view returns(uint256)
func (_Vnft *VnftCaller) GetEmptyNftCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getEmptyNftCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEmptyNftCounts is a free data retrieval call binding the contract method 0x0ec3be08.
//
// Solidity: function getEmptyNftCounts() view returns(uint256)
func (_Vnft *VnftSession) GetEmptyNftCounts() (*big.Int, error) {
	return _Vnft.Contract.GetEmptyNftCounts(&_Vnft.CallOpts)
}

// GetEmptyNftCounts is a free data retrieval call binding the contract method 0x0ec3be08.
//
// Solidity: function getEmptyNftCounts() view returns(uint256)
func (_Vnft *VnftCallerSession) GetEmptyNftCounts() (*big.Int, error) {
	return _Vnft.Contract.GetEmptyNftCounts(&_Vnft.CallOpts)
}

// GetEmptyNftCountsOfOperator is a free data retrieval call binding the contract method 0x1a85b068.
//
// Solidity: function getEmptyNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCaller) GetEmptyNftCountsOfOperator(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getEmptyNftCountsOfOperator", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEmptyNftCountsOfOperator is a free data retrieval call binding the contract method 0x1a85b068.
//
// Solidity: function getEmptyNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftSession) GetEmptyNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetEmptyNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetEmptyNftCountsOfOperator is a free data retrieval call binding the contract method 0x1a85b068.
//
// Solidity: function getEmptyNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCallerSession) GetEmptyNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetEmptyNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetNextValidatorWithdrawalCredential is a free data retrieval call binding the contract method 0xbf1208d7.
//
// Solidity: function getNextValidatorWithdrawalCredential(uint256 _operatorId) view returns(bytes)
func (_Vnft *VnftCaller) GetNextValidatorWithdrawalCredential(opts *bind.CallOpts, _operatorId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getNextValidatorWithdrawalCredential", _operatorId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNextValidatorWithdrawalCredential is a free data retrieval call binding the contract method 0xbf1208d7.
//
// Solidity: function getNextValidatorWithdrawalCredential(uint256 _operatorId) view returns(bytes)
func (_Vnft *VnftSession) GetNextValidatorWithdrawalCredential(_operatorId *big.Int) ([]byte, error) {
	return _Vnft.Contract.GetNextValidatorWithdrawalCredential(&_Vnft.CallOpts, _operatorId)
}

// GetNextValidatorWithdrawalCredential is a free data retrieval call binding the contract method 0xbf1208d7.
//
// Solidity: function getNextValidatorWithdrawalCredential(uint256 _operatorId) view returns(bytes)
func (_Vnft *VnftCallerSession) GetNextValidatorWithdrawalCredential(_operatorId *big.Int) ([]byte, error) {
	return _Vnft.Contract.GetNextValidatorWithdrawalCredential(&_Vnft.CallOpts, _operatorId)
}

// GetNftExitBlockNumbers is a free data retrieval call binding the contract method 0x79ff5d1d.
//
// Solidity: function getNftExitBlockNumbers(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftCaller) GetNftExitBlockNumbers(opts *bind.CallOpts, _tokenIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getNftExitBlockNumbers", _tokenIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNftExitBlockNumbers is a free data retrieval call binding the contract method 0x79ff5d1d.
//
// Solidity: function getNftExitBlockNumbers(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftSession) GetNftExitBlockNumbers(_tokenIds []*big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.GetNftExitBlockNumbers(&_Vnft.CallOpts, _tokenIds)
}

// GetNftExitBlockNumbers is a free data retrieval call binding the contract method 0x79ff5d1d.
//
// Solidity: function getNftExitBlockNumbers(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftCallerSession) GetNftExitBlockNumbers(_tokenIds []*big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.GetNftExitBlockNumbers(&_Vnft.CallOpts, _tokenIds)
}

// GetTotalActiveNftCounts is a free data retrieval call binding the contract method 0xebbecb93.
//
// Solidity: function getTotalActiveNftCounts() view returns(uint256)
func (_Vnft *VnftCaller) GetTotalActiveNftCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getTotalActiveNftCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveNftCounts is a free data retrieval call binding the contract method 0xebbecb93.
//
// Solidity: function getTotalActiveNftCounts() view returns(uint256)
func (_Vnft *VnftSession) GetTotalActiveNftCounts() (*big.Int, error) {
	return _Vnft.Contract.GetTotalActiveNftCounts(&_Vnft.CallOpts)
}

// GetTotalActiveNftCounts is a free data retrieval call binding the contract method 0xebbecb93.
//
// Solidity: function getTotalActiveNftCounts() view returns(uint256)
func (_Vnft *VnftCallerSession) GetTotalActiveNftCounts() (*big.Int, error) {
	return _Vnft.Contract.GetTotalActiveNftCounts(&_Vnft.CallOpts)
}

// GetUserActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x1cd3767b.
//
// Solidity: function getUserActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCaller) GetUserActiveNftCountsOfOperator(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getUserActiveNftCountsOfOperator", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x1cd3767b.
//
// Solidity: function getUserActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftSession) GetUserActiveNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetUserActiveNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetUserActiveNftCountsOfOperator is a free data retrieval call binding the contract method 0x1cd3767b.
//
// Solidity: function getUserActiveNftCountsOfOperator(uint256 _operatorId) view returns(uint256)
func (_Vnft *VnftCallerSession) GetUserActiveNftCountsOfOperator(_operatorId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.GetUserActiveNftCountsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// GetUserNftGasHeight is a free data retrieval call binding the contract method 0x94d93b0b.
//
// Solidity: function getUserNftGasHeight(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftCaller) GetUserNftGasHeight(opts *bind.CallOpts, _tokenIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getUserNftGasHeight", _tokenIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserNftGasHeight is a free data retrieval call binding the contract method 0x94d93b0b.
//
// Solidity: function getUserNftGasHeight(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftSession) GetUserNftGasHeight(_tokenIds []*big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.GetUserNftGasHeight(&_Vnft.CallOpts, _tokenIds)
}

// GetUserNftGasHeight is a free data retrieval call binding the contract method 0x94d93b0b.
//
// Solidity: function getUserNftGasHeight(uint256[] _tokenIds) view returns(uint256[])
func (_Vnft *VnftCallerSession) GetUserNftGasHeight(_tokenIds []*big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.GetUserNftGasHeight(&_Vnft.CallOpts, _tokenIds)
}

// GetUserNftWithdrawalCredentialOfTokenId is a free data retrieval call binding the contract method 0x25171570.
//
// Solidity: function getUserNftWithdrawalCredentialOfTokenId(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftCaller) GetUserNftWithdrawalCredentialOfTokenId(opts *bind.CallOpts, _tokenId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "getUserNftWithdrawalCredentialOfTokenId", _tokenId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetUserNftWithdrawalCredentialOfTokenId is a free data retrieval call binding the contract method 0x25171570.
//
// Solidity: function getUserNftWithdrawalCredentialOfTokenId(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftSession) GetUserNftWithdrawalCredentialOfTokenId(_tokenId *big.Int) ([]byte, error) {
	return _Vnft.Contract.GetUserNftWithdrawalCredentialOfTokenId(&_Vnft.CallOpts, _tokenId)
}

// GetUserNftWithdrawalCredentialOfTokenId is a free data retrieval call binding the contract method 0x25171570.
//
// Solidity: function getUserNftWithdrawalCredentialOfTokenId(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftCallerSession) GetUserNftWithdrawalCredentialOfTokenId(_tokenId *big.Int) ([]byte, error) {
	return _Vnft.Contract.GetUserNftWithdrawalCredentialOfTokenId(&_Vnft.CallOpts, _tokenId)
}

// InitHeightOf is a free data retrieval call binding the contract method 0x95ef3d40.
//
// Solidity: function initHeightOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftCaller) InitHeightOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "initHeightOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitHeightOf is a free data retrieval call binding the contract method 0x95ef3d40.
//
// Solidity: function initHeightOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftSession) InitHeightOf(_tokenId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.InitHeightOf(&_Vnft.CallOpts, _tokenId)
}

// InitHeightOf is a free data retrieval call binding the contract method 0x95ef3d40.
//
// Solidity: function initHeightOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftCallerSession) InitHeightOf(_tokenId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.InitHeightOf(&_Vnft.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Vnft *VnftCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Vnft *VnftSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Vnft.Contract.IsApprovedForAll(&_Vnft.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Vnft *VnftCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Vnft.Contract.IsApprovedForAll(&_Vnft.CallOpts, _owner, _operator)
}

// LastOwnerOf is a free data retrieval call binding the contract method 0xa2ddd349.
//
// Solidity: function lastOwnerOf(uint256 _tokenId) view returns(address)
func (_Vnft *VnftCaller) LastOwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "lastOwnerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastOwnerOf is a free data retrieval call binding the contract method 0xa2ddd349.
//
// Solidity: function lastOwnerOf(uint256 _tokenId) view returns(address)
func (_Vnft *VnftSession) LastOwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.LastOwnerOf(&_Vnft.CallOpts, _tokenId)
}

// LastOwnerOf is a free data retrieval call binding the contract method 0xa2ddd349.
//
// Solidity: function lastOwnerOf(uint256 _tokenId) view returns(address)
func (_Vnft *VnftCallerSession) LastOwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.LastOwnerOf(&_Vnft.CallOpts, _tokenId)
}

// LastOwners is a free data retrieval call binding the contract method 0xd09fe08f.
//
// Solidity: function lastOwners(uint256 ) view returns(address)
func (_Vnft *VnftCaller) LastOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "lastOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastOwners is a free data retrieval call binding the contract method 0xd09fe08f.
//
// Solidity: function lastOwners(uint256 ) view returns(address)
func (_Vnft *VnftSession) LastOwners(arg0 *big.Int) (common.Address, error) {
	return _Vnft.Contract.LastOwners(&_Vnft.CallOpts, arg0)
}

// LastOwners is a free data retrieval call binding the contract method 0xd09fe08f.
//
// Solidity: function lastOwners(uint256 ) view returns(address)
func (_Vnft *VnftCallerSession) LastOwners(arg0 *big.Int) (common.Address, error) {
	return _Vnft.Contract.LastOwners(&_Vnft.CallOpts, arg0)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Vnft *VnftCaller) LiquidStakingContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "liquidStakingContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Vnft *VnftSession) LiquidStakingContractAddress() (common.Address, error) {
	return _Vnft.Contract.LiquidStakingContractAddress(&_Vnft.CallOpts)
}

// LiquidStakingContractAddress is a free data retrieval call binding the contract method 0x6404a4c7.
//
// Solidity: function liquidStakingContractAddress() view returns(address)
func (_Vnft *VnftCallerSession) LiquidStakingContractAddress() (common.Address, error) {
	return _Vnft.Contract.LiquidStakingContractAddress(&_Vnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vnft *VnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vnft *VnftSession) Name() (string, error) {
	return _Vnft.Contract.Name(&_Vnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vnft *VnftCallerSession) Name() (string, error) {
	return _Vnft.Contract.Name(&_Vnft.CallOpts)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address _owner) view returns(uint256)
func (_Vnft *VnftCaller) NumberMinted(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "numberMinted", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address _owner) view returns(uint256)
func (_Vnft *VnftSession) NumberMinted(_owner common.Address) (*big.Int, error) {
	return _Vnft.Contract.NumberMinted(&_Vnft.CallOpts, _owner)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address _owner) view returns(uint256)
func (_Vnft *VnftCallerSession) NumberMinted(_owner common.Address) (*big.Int, error) {
	return _Vnft.Contract.NumberMinted(&_Vnft.CallOpts, _owner)
}

// OperatorEmptyNftIndex is a free data retrieval call binding the contract method 0xa4cd56bb.
//
// Solidity: function operatorEmptyNftIndex(uint256 ) view returns(uint256)
func (_Vnft *VnftCaller) OperatorEmptyNftIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "operatorEmptyNftIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorEmptyNftIndex is a free data retrieval call binding the contract method 0xa4cd56bb.
//
// Solidity: function operatorEmptyNftIndex(uint256 ) view returns(uint256)
func (_Vnft *VnftSession) OperatorEmptyNftIndex(arg0 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorEmptyNftIndex(&_Vnft.CallOpts, arg0)
}

// OperatorEmptyNftIndex is a free data retrieval call binding the contract method 0xa4cd56bb.
//
// Solidity: function operatorEmptyNftIndex(uint256 ) view returns(uint256)
func (_Vnft *VnftCallerSession) OperatorEmptyNftIndex(arg0 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorEmptyNftIndex(&_Vnft.CallOpts, arg0)
}

// OperatorEmptyNfts is a free data retrieval call binding the contract method 0xd127808c.
//
// Solidity: function operatorEmptyNfts(uint256 , uint256 ) view returns(uint256)
func (_Vnft *VnftCaller) OperatorEmptyNfts(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "operatorEmptyNfts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorEmptyNfts is a free data retrieval call binding the contract method 0xd127808c.
//
// Solidity: function operatorEmptyNfts(uint256 , uint256 ) view returns(uint256)
func (_Vnft *VnftSession) OperatorEmptyNfts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorEmptyNfts(&_Vnft.CallOpts, arg0, arg1)
}

// OperatorEmptyNfts is a free data retrieval call binding the contract method 0xd127808c.
//
// Solidity: function operatorEmptyNfts(uint256 , uint256 ) view returns(uint256)
func (_Vnft *VnftCallerSession) OperatorEmptyNfts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorEmptyNfts(&_Vnft.CallOpts, arg0, arg1)
}

// OperatorOf is a free data retrieval call binding the contract method 0xf41875e1.
//
// Solidity: function operatorOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftCaller) OperatorOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "operatorOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorOf is a free data retrieval call binding the contract method 0xf41875e1.
//
// Solidity: function operatorOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftSession) OperatorOf(_tokenId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorOf(&_Vnft.CallOpts, _tokenId)
}

// OperatorOf is a free data retrieval call binding the contract method 0xf41875e1.
//
// Solidity: function operatorOf(uint256 _tokenId) view returns(uint256)
func (_Vnft *VnftCallerSession) OperatorOf(_tokenId *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorOf(&_Vnft.CallOpts, _tokenId)
}

// OperatorRecords is a free data retrieval call binding the contract method 0x2a6e5e37.
//
// Solidity: function operatorRecords(uint256 ) view returns(uint256)
func (_Vnft *VnftCaller) OperatorRecords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "operatorRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorRecords is a free data retrieval call binding the contract method 0x2a6e5e37.
//
// Solidity: function operatorRecords(uint256 ) view returns(uint256)
func (_Vnft *VnftSession) OperatorRecords(arg0 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorRecords(&_Vnft.CallOpts, arg0)
}

// OperatorRecords is a free data retrieval call binding the contract method 0x2a6e5e37.
//
// Solidity: function operatorRecords(uint256 ) view returns(uint256)
func (_Vnft *VnftCallerSession) OperatorRecords(arg0 *big.Int) (*big.Int, error) {
	return _Vnft.Contract.OperatorRecords(&_Vnft.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vnft *VnftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vnft *VnftSession) Owner() (common.Address, error) {
	return _Vnft.Contract.Owner(&_Vnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vnft *VnftCallerSession) Owner() (common.Address, error) {
	return _Vnft.Contract.Owner(&_Vnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vnft *VnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vnft *VnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.OwnerOf(&_Vnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vnft *VnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Vnft.Contract.OwnerOf(&_Vnft.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vnft *VnftCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vnft *VnftSession) ProxiableUUID() ([32]byte, error) {
	return _Vnft.Contract.ProxiableUUID(&_Vnft.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vnft *VnftCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Vnft.Contract.ProxiableUUID(&_Vnft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vnft *VnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vnft *VnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vnft.Contract.SupportsInterface(&_Vnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vnft *VnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vnft.Contract.SupportsInterface(&_Vnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vnft *VnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vnft *VnftSession) Symbol() (string, error) {
	return _Vnft.Contract.Symbol(&_Vnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vnft *VnftCallerSession) Symbol() (string, error) {
	return _Vnft.Contract.Symbol(&_Vnft.CallOpts)
}

// TokenOfValidator is a free data retrieval call binding the contract method 0x3d09fdc5.
//
// Solidity: function tokenOfValidator(bytes _pubkey) view returns(uint256)
func (_Vnft *VnftCaller) TokenOfValidator(opts *bind.CallOpts, _pubkey []byte) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "tokenOfValidator", _pubkey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfValidator is a free data retrieval call binding the contract method 0x3d09fdc5.
//
// Solidity: function tokenOfValidator(bytes _pubkey) view returns(uint256)
func (_Vnft *VnftSession) TokenOfValidator(_pubkey []byte) (*big.Int, error) {
	return _Vnft.Contract.TokenOfValidator(&_Vnft.CallOpts, _pubkey)
}

// TokenOfValidator is a free data retrieval call binding the contract method 0x3d09fdc5.
//
// Solidity: function tokenOfValidator(bytes _pubkey) view returns(uint256)
func (_Vnft *VnftCallerSession) TokenOfValidator(_pubkey []byte) (*big.Int, error) {
	return _Vnft.Contract.TokenOfValidator(&_Vnft.CallOpts, _pubkey)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vnft *VnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vnft *VnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Vnft.Contract.TokenURI(&_Vnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vnft *VnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Vnft.Contract.TokenURI(&_Vnft.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Vnft *VnftCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Vnft *VnftSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Vnft.Contract.TokensOfOwner(&_Vnft.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Vnft *VnftCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Vnft.Contract.TokensOfOwner(&_Vnft.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Vnft *VnftCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Vnft *VnftSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.TokensOfOwnerIn(&_Vnft.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Vnft *VnftCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Vnft.Contract.TokensOfOwnerIn(&_Vnft.CallOpts, owner, start, stop)
}

// TotalExitButNoBurnNftCounts is a free data retrieval call binding the contract method 0x1fdb58b4.
//
// Solidity: function totalExitButNoBurnNftCounts() view returns(uint256)
func (_Vnft *VnftCaller) TotalExitButNoBurnNftCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "totalExitButNoBurnNftCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalExitButNoBurnNftCounts is a free data retrieval call binding the contract method 0x1fdb58b4.
//
// Solidity: function totalExitButNoBurnNftCounts() view returns(uint256)
func (_Vnft *VnftSession) TotalExitButNoBurnNftCounts() (*big.Int, error) {
	return _Vnft.Contract.TotalExitButNoBurnNftCounts(&_Vnft.CallOpts)
}

// TotalExitButNoBurnNftCounts is a free data retrieval call binding the contract method 0x1fdb58b4.
//
// Solidity: function totalExitButNoBurnNftCounts() view returns(uint256)
func (_Vnft *VnftCallerSession) TotalExitButNoBurnNftCounts() (*big.Int, error) {
	return _Vnft.Contract.TotalExitButNoBurnNftCounts(&_Vnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vnft *VnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vnft *VnftSession) TotalSupply() (*big.Int, error) {
	return _Vnft.Contract.TotalSupply(&_Vnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vnft *VnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Vnft.Contract.TotalSupply(&_Vnft.CallOpts)
}

// ValidatorExists is a free data retrieval call binding the contract method 0x1d06b1cd.
//
// Solidity: function validatorExists(bytes _pubkey) view returns(bool)
func (_Vnft *VnftCaller) ValidatorExists(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validatorExists", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorExists is a free data retrieval call binding the contract method 0x1d06b1cd.
//
// Solidity: function validatorExists(bytes _pubkey) view returns(bool)
func (_Vnft *VnftSession) ValidatorExists(_pubkey []byte) (bool, error) {
	return _Vnft.Contract.ValidatorExists(&_Vnft.CallOpts, _pubkey)
}

// ValidatorExists is a free data retrieval call binding the contract method 0x1d06b1cd.
//
// Solidity: function validatorExists(bytes _pubkey) view returns(bool)
func (_Vnft *VnftCallerSession) ValidatorExists(_pubkey []byte) (bool, error) {
	return _Vnft.Contract.ValidatorExists(&_Vnft.CallOpts, _pubkey)
}

// ValidatorOf is a free data retrieval call binding the contract method 0x5c674ff4.
//
// Solidity: function validatorOf(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftCaller) ValidatorOf(opts *bind.CallOpts, _tokenId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validatorOf", _tokenId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidatorOf is a free data retrieval call binding the contract method 0x5c674ff4.
//
// Solidity: function validatorOf(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftSession) ValidatorOf(_tokenId *big.Int) ([]byte, error) {
	return _Vnft.Contract.ValidatorOf(&_Vnft.CallOpts, _tokenId)
}

// ValidatorOf is a free data retrieval call binding the contract method 0x5c674ff4.
//
// Solidity: function validatorOf(uint256 _tokenId) view returns(bytes)
func (_Vnft *VnftCallerSession) ValidatorOf(_tokenId *big.Int) ([]byte, error) {
	return _Vnft.Contract.ValidatorOf(&_Vnft.CallOpts, _tokenId)
}

// ValidatorRecords is a free data retrieval call binding the contract method 0x052bbca0.
//
// Solidity: function validatorRecords(bytes ) view returns(uint256)
func (_Vnft *VnftCaller) ValidatorRecords(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validatorRecords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRecords is a free data retrieval call binding the contract method 0x052bbca0.
//
// Solidity: function validatorRecords(bytes ) view returns(uint256)
func (_Vnft *VnftSession) ValidatorRecords(arg0 []byte) (*big.Int, error) {
	return _Vnft.Contract.ValidatorRecords(&_Vnft.CallOpts, arg0)
}

// ValidatorRecords is a free data retrieval call binding the contract method 0x052bbca0.
//
// Solidity: function validatorRecords(bytes ) view returns(uint256)
func (_Vnft *VnftCallerSession) ValidatorRecords(arg0 []byte) (*big.Int, error) {
	return _Vnft.Contract.ValidatorRecords(&_Vnft.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 operatorId, uint256 initHeight, bytes pubkey)
func (_Vnft *VnftCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OperatorId *big.Int
	InitHeight *big.Int
	Pubkey     []byte
}, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		OperatorId *big.Int
		InitHeight *big.Int
		Pubkey     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InitHeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pubkey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 operatorId, uint256 initHeight, bytes pubkey)
func (_Vnft *VnftSession) Validators(arg0 *big.Int) (struct {
	OperatorId *big.Int
	InitHeight *big.Int
	Pubkey     []byte
}, error) {
	return _Vnft.Contract.Validators(&_Vnft.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 operatorId, uint256 initHeight, bytes pubkey)
func (_Vnft *VnftCallerSession) Validators(arg0 *big.Int) (struct {
	OperatorId *big.Int
	InitHeight *big.Int
	Pubkey     []byte
}, error) {
	return _Vnft.Contract.Validators(&_Vnft.CallOpts, arg0)
}

// ValidatorsOfOperator is a free data retrieval call binding the contract method 0x76abbc34.
//
// Solidity: function validatorsOfOperator(uint256 _operatorId) view returns(bytes[])
func (_Vnft *VnftCaller) ValidatorsOfOperator(opts *bind.CallOpts, _operatorId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validatorsOfOperator", _operatorId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ValidatorsOfOperator is a free data retrieval call binding the contract method 0x76abbc34.
//
// Solidity: function validatorsOfOperator(uint256 _operatorId) view returns(bytes[])
func (_Vnft *VnftSession) ValidatorsOfOperator(_operatorId *big.Int) ([][]byte, error) {
	return _Vnft.Contract.ValidatorsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// ValidatorsOfOperator is a free data retrieval call binding the contract method 0x76abbc34.
//
// Solidity: function validatorsOfOperator(uint256 _operatorId) view returns(bytes[])
func (_Vnft *VnftCallerSession) ValidatorsOfOperator(_operatorId *big.Int) ([][]byte, error) {
	return _Vnft.Contract.ValidatorsOfOperator(&_Vnft.CallOpts, _operatorId)
}

// ValidatorsOfOwner is a free data retrieval call binding the contract method 0xa25ab6d4.
//
// Solidity: function validatorsOfOwner(address _owner) view returns(bytes[])
func (_Vnft *VnftCaller) ValidatorsOfOwner(opts *bind.CallOpts, _owner common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Vnft.contract.Call(opts, &out, "validatorsOfOwner", _owner)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ValidatorsOfOwner is a free data retrieval call binding the contract method 0xa25ab6d4.
//
// Solidity: function validatorsOfOwner(address _owner) view returns(bytes[])
func (_Vnft *VnftSession) ValidatorsOfOwner(_owner common.Address) ([][]byte, error) {
	return _Vnft.Contract.ValidatorsOfOwner(&_Vnft.CallOpts, _owner)
}

// ValidatorsOfOwner is a free data retrieval call binding the contract method 0xa25ab6d4.
//
// Solidity: function validatorsOfOwner(address _owner) view returns(bytes[])
func (_Vnft *VnftCallerSession) ValidatorsOfOwner(_owner common.Address) ([][]byte, error) {
	return _Vnft.Contract.ValidatorsOfOwner(&_Vnft.CallOpts, _owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Vnft *VnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.Approve(&_Vnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.Approve(&_Vnft.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vnft *VnftTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vnft *VnftSession) Initialize() (*types.Transaction, error) {
	return _Vnft.Contract.Initialize(&_Vnft.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vnft *VnftTransactorSession) Initialize() (*types.Transaction, error) {
	return _Vnft.Contract.Initialize(&_Vnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vnft *VnftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vnft *VnftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vnft.Contract.RenounceOwnership(&_Vnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vnft *VnftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vnft.Contract.RenounceOwnership(&_Vnft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SafeTransferFrom(&_Vnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SafeTransferFrom(&_Vnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Vnft *VnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Vnft *VnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Vnft.Contract.SafeTransferFrom0(&_Vnft.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Vnft *VnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Vnft.Contract.SafeTransferFrom0(&_Vnft.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vnft *VnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vnft *VnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vnft.Contract.SetApprovalForAll(&_Vnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vnft *VnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vnft.Contract.SetApprovalForAll(&_Vnft.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Vnft *VnftTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Vnft *VnftSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Vnft.Contract.SetBaseURI(&_Vnft.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Vnft *VnftTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Vnft.Contract.SetBaseURI(&_Vnft.TransactOpts, _baseURI)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Vnft *VnftTransactor) SetLiquidStaking(opts *bind.TransactOpts, _liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "setLiquidStaking", _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Vnft *VnftSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.SetLiquidStaking(&_Vnft.TransactOpts, _liquidStakingContractAddress)
}

// SetLiquidStaking is a paid mutator transaction binding the contract method 0x08211be5.
//
// Solidity: function setLiquidStaking(address _liquidStakingContractAddress) returns()
func (_Vnft *VnftTransactorSession) SetLiquidStaking(_liquidStakingContractAddress common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.SetLiquidStaking(&_Vnft.TransactOpts, _liquidStakingContractAddress)
}

// SetNftExitBlockNumbers is a paid mutator transaction binding the contract method 0x677f068a.
//
// Solidity: function setNftExitBlockNumbers(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Vnft *VnftTransactor) SetNftExitBlockNumbers(opts *bind.TransactOpts, _tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "setNftExitBlockNumbers", _tokenIds, _exitBlockNumbers)
}

// SetNftExitBlockNumbers is a paid mutator transaction binding the contract method 0x677f068a.
//
// Solidity: function setNftExitBlockNumbers(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Vnft *VnftSession) SetNftExitBlockNumbers(_tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SetNftExitBlockNumbers(&_Vnft.TransactOpts, _tokenIds, _exitBlockNumbers)
}

// SetNftExitBlockNumbers is a paid mutator transaction binding the contract method 0x677f068a.
//
// Solidity: function setNftExitBlockNumbers(uint256[] _tokenIds, uint256[] _exitBlockNumbers) returns()
func (_Vnft *VnftTransactorSession) SetNftExitBlockNumbers(_tokenIds []*big.Int, _exitBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SetNftExitBlockNumbers(&_Vnft.TransactOpts, _tokenIds, _exitBlockNumbers)
}

// SetUserNftGasHeight is a paid mutator transaction binding the contract method 0xba4b4cd4.
//
// Solidity: function setUserNftGasHeight(uint256 _tokenId, uint256 _number) returns()
func (_Vnft *VnftTransactor) SetUserNftGasHeight(opts *bind.TransactOpts, _tokenId *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "setUserNftGasHeight", _tokenId, _number)
}

// SetUserNftGasHeight is a paid mutator transaction binding the contract method 0xba4b4cd4.
//
// Solidity: function setUserNftGasHeight(uint256 _tokenId, uint256 _number) returns()
func (_Vnft *VnftSession) SetUserNftGasHeight(_tokenId *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SetUserNftGasHeight(&_Vnft.TransactOpts, _tokenId, _number)
}

// SetUserNftGasHeight is a paid mutator transaction binding the contract method 0xba4b4cd4.
//
// Solidity: function setUserNftGasHeight(uint256 _tokenId, uint256 _number) returns()
func (_Vnft *VnftTransactorSession) SetUserNftGasHeight(_tokenId *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.SetUserNftGasHeight(&_Vnft.TransactOpts, _tokenId, _number)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.TransferFrom(&_Vnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Vnft *VnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.TransferFrom(&_Vnft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vnft *VnftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vnft *VnftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.TransferOwnership(&_Vnft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vnft *VnftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.TransferOwnership(&_Vnft.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vnft *VnftTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vnft *VnftSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.UpgradeTo(&_Vnft.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vnft *VnftTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vnft.Contract.UpgradeTo(&_Vnft.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vnft *VnftTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vnft *VnftSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vnft.Contract.UpgradeToAndCall(&_Vnft.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vnft *VnftTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vnft.Contract.UpgradeToAndCall(&_Vnft.TransactOpts, newImplementation, data)
}

// WhiteListBurn is a paid mutator transaction binding the contract method 0xc2c3a60d.
//
// Solidity: function whiteListBurn(uint256 _tokenId) returns()
func (_Vnft *VnftTransactor) WhiteListBurn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "whiteListBurn", _tokenId)
}

// WhiteListBurn is a paid mutator transaction binding the contract method 0xc2c3a60d.
//
// Solidity: function whiteListBurn(uint256 _tokenId) returns()
func (_Vnft *VnftSession) WhiteListBurn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.WhiteListBurn(&_Vnft.TransactOpts, _tokenId)
}

// WhiteListBurn is a paid mutator transaction binding the contract method 0xc2c3a60d.
//
// Solidity: function whiteListBurn(uint256 _tokenId) returns()
func (_Vnft *VnftTransactorSession) WhiteListBurn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.WhiteListBurn(&_Vnft.TransactOpts, _tokenId)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x0314cb05.
//
// Solidity: function whiteListMint(bytes _pubkey, bytes _withdrawalCredentials, address _to, uint256 _operatorId) returns(uint256)
func (_Vnft *VnftTransactor) WhiteListMint(opts *bind.TransactOpts, _pubkey []byte, _withdrawalCredentials []byte, _to common.Address, _operatorId *big.Int) (*types.Transaction, error) {
	return _Vnft.contract.Transact(opts, "whiteListMint", _pubkey, _withdrawalCredentials, _to, _operatorId)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x0314cb05.
//
// Solidity: function whiteListMint(bytes _pubkey, bytes _withdrawalCredentials, address _to, uint256 _operatorId) returns(uint256)
func (_Vnft *VnftSession) WhiteListMint(_pubkey []byte, _withdrawalCredentials []byte, _to common.Address, _operatorId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.WhiteListMint(&_Vnft.TransactOpts, _pubkey, _withdrawalCredentials, _to, _operatorId)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x0314cb05.
//
// Solidity: function whiteListMint(bytes _pubkey, bytes _withdrawalCredentials, address _to, uint256 _operatorId) returns(uint256)
func (_Vnft *VnftTransactorSession) WhiteListMint(_pubkey []byte, _withdrawalCredentials []byte, _to common.Address, _operatorId *big.Int) (*types.Transaction, error) {
	return _Vnft.Contract.WhiteListMint(&_Vnft.TransactOpts, _pubkey, _withdrawalCredentials, _to, _operatorId)
}

// VnftAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Vnft contract.
type VnftAdminChangedIterator struct {
	Event *VnftAdminChanged // Event containing the contract specifics and raw log

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
func (it *VnftAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftAdminChanged)
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
		it.Event = new(VnftAdminChanged)
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
func (it *VnftAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftAdminChanged represents a AdminChanged event raised by the Vnft contract.
type VnftAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vnft *VnftFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VnftAdminChangedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VnftAdminChangedIterator{contract: _Vnft.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vnft *VnftFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VnftAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftAdminChanged)
				if err := _Vnft.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vnft *VnftFilterer) ParseAdminChanged(log types.Log) (*VnftAdminChanged, error) {
	event := new(VnftAdminChanged)
	if err := _Vnft.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vnft contract.
type VnftApprovalIterator struct {
	Event *VnftApproval // Event containing the contract specifics and raw log

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
func (it *VnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftApproval)
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
		it.Event = new(VnftApproval)
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
func (it *VnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftApproval represents a Approval event raised by the Vnft contract.
type VnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Vnft *VnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VnftApprovalIterator, error) {

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

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VnftApprovalIterator{contract: _Vnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Vnft *VnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftApproval)
				if err := _Vnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Vnft *VnftFilterer) ParseApproval(log types.Log) (*VnftApproval, error) {
	event := new(VnftApproval)
	if err := _Vnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Vnft contract.
type VnftApprovalForAllIterator struct {
	Event *VnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftApprovalForAll)
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
		it.Event = new(VnftApprovalForAll)
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
func (it *VnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftApprovalForAll represents a ApprovalForAll event raised by the Vnft contract.
type VnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Vnft *VnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VnftApprovalForAllIterator{contract: _Vnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Vnft *VnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftApprovalForAll)
				if err := _Vnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Vnft *VnftFilterer) ParseApprovalForAll(log types.Log) (*VnftApprovalForAll, error) {
	event := new(VnftApprovalForAll)
	if err := _Vnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the Vnft contract.
type VnftBaseURIChangedIterator struct {
	Event *VnftBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *VnftBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftBaseURIChanged)
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
		it.Event = new(VnftBaseURIChanged)
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
func (it *VnftBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftBaseURIChanged represents a BaseURIChanged event raised by the Vnft contract.
type VnftBaseURIChanged struct {
	Before string
	After  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0xc41b7cb64e5be01af4afc2641afc861432136270f4206b7773f229b658b96699.
//
// Solidity: event BaseURIChanged(string _before, string _after)
func (_Vnft *VnftFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*VnftBaseURIChangedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &VnftBaseURIChangedIterator{contract: _Vnft.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0xc41b7cb64e5be01af4afc2641afc861432136270f4206b7773f229b658b96699.
//
// Solidity: event BaseURIChanged(string _before, string _after)
func (_Vnft *VnftFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *VnftBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftBaseURIChanged)
				if err := _Vnft.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0xc41b7cb64e5be01af4afc2641afc861432136270f4206b7773f229b658b96699.
//
// Solidity: event BaseURIChanged(string _before, string _after)
func (_Vnft *VnftFilterer) ParseBaseURIChanged(log types.Log) (*VnftBaseURIChanged, error) {
	event := new(VnftBaseURIChanged)
	if err := _Vnft.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Vnft contract.
type VnftBeaconUpgradedIterator struct {
	Event *VnftBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VnftBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftBeaconUpgraded)
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
		it.Event = new(VnftBeaconUpgraded)
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
func (it *VnftBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftBeaconUpgraded represents a BeaconUpgraded event raised by the Vnft contract.
type VnftBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vnft *VnftFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VnftBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VnftBeaconUpgradedIterator{contract: _Vnft.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vnft *VnftFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VnftBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftBeaconUpgraded)
				if err := _Vnft.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vnft *VnftFilterer) ParseBeaconUpgraded(log types.Log) (*VnftBeaconUpgraded, error) {
	event := new(VnftBeaconUpgraded)
	if err := _Vnft.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Vnft contract.
type VnftConsecutiveTransferIterator struct {
	Event *VnftConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *VnftConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftConsecutiveTransfer)
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
		it.Event = new(VnftConsecutiveTransfer)
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
func (it *VnftConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Vnft contract.
type VnftConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Vnft *VnftFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*VnftConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VnftConsecutiveTransferIterator{contract: _Vnft.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Vnft *VnftFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *VnftConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftConsecutiveTransfer)
				if err := _Vnft.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Vnft *VnftFilterer) ParseConsecutiveTransfer(log types.Log) (*VnftConsecutiveTransfer, error) {
	event := new(VnftConsecutiveTransfer)
	if err := _Vnft.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vnft contract.
type VnftInitializedIterator struct {
	Event *VnftInitialized // Event containing the contract specifics and raw log

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
func (it *VnftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftInitialized)
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
		it.Event = new(VnftInitialized)
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
func (it *VnftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftInitialized represents a Initialized event raised by the Vnft contract.
type VnftInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vnft *VnftFilterer) FilterInitialized(opts *bind.FilterOpts) (*VnftInitializedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VnftInitializedIterator{contract: _Vnft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vnft *VnftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VnftInitialized) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftInitialized)
				if err := _Vnft.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vnft *VnftFilterer) ParseInitialized(log types.Log) (*VnftInitialized, error) {
	event := new(VnftInitialized)
	if err := _Vnft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftLiquidStakingChangedIterator is returned from FilterLiquidStakingChanged and is used to iterate over the raw logs and unpacked data for LiquidStakingChanged events raised by the Vnft contract.
type VnftLiquidStakingChangedIterator struct {
	Event *VnftLiquidStakingChanged // Event containing the contract specifics and raw log

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
func (it *VnftLiquidStakingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftLiquidStakingChanged)
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
		it.Event = new(VnftLiquidStakingChanged)
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
func (it *VnftLiquidStakingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftLiquidStakingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftLiquidStakingChanged represents a LiquidStakingChanged event raised by the Vnft contract.
type VnftLiquidStakingChanged struct {
	Before common.Address
	After  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingChanged is a free log retrieval operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Vnft *VnftFilterer) FilterLiquidStakingChanged(opts *bind.FilterOpts) (*VnftLiquidStakingChangedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return &VnftLiquidStakingChangedIterator{contract: _Vnft.contract, event: "LiquidStakingChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingChanged is a free log subscription operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Vnft *VnftFilterer) WatchLiquidStakingChanged(opts *bind.WatchOpts, sink chan<- *VnftLiquidStakingChanged) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "LiquidStakingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftLiquidStakingChanged)
				if err := _Vnft.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
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

// ParseLiquidStakingChanged is a log parse operation binding the contract event 0x516e369f70685d2905d214a9e8567098c02a0e00f5a55bd30baca6b51d446cef.
//
// Solidity: event LiquidStakingChanged(address _before, address _after)
func (_Vnft *VnftFilterer) ParseLiquidStakingChanged(log types.Log) (*VnftLiquidStakingChanged, error) {
	event := new(VnftLiquidStakingChanged)
	if err := _Vnft.contract.UnpackLog(event, "LiquidStakingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftNFTBurnedIterator is returned from FilterNFTBurned and is used to iterate over the raw logs and unpacked data for NFTBurned events raised by the Vnft contract.
type VnftNFTBurnedIterator struct {
	Event *VnftNFTBurned // Event containing the contract specifics and raw log

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
func (it *VnftNFTBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftNFTBurned)
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
		it.Event = new(VnftNFTBurned)
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
func (it *VnftNFTBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftNFTBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftNFTBurned represents a NFTBurned event raised by the Vnft contract.
type VnftNFTBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTBurned is a free log retrieval operation binding the contract event 0x3c176691ca154a2f6fe978a2a633a33ee77dbe2902e67a75400720845a4b2ce1.
//
// Solidity: event NFTBurned(uint256 _tokenId)
func (_Vnft *VnftFilterer) FilterNFTBurned(opts *bind.FilterOpts) (*VnftNFTBurnedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "NFTBurned")
	if err != nil {
		return nil, err
	}
	return &VnftNFTBurnedIterator{contract: _Vnft.contract, event: "NFTBurned", logs: logs, sub: sub}, nil
}

// WatchNFTBurned is a free log subscription operation binding the contract event 0x3c176691ca154a2f6fe978a2a633a33ee77dbe2902e67a75400720845a4b2ce1.
//
// Solidity: event NFTBurned(uint256 _tokenId)
func (_Vnft *VnftFilterer) WatchNFTBurned(opts *bind.WatchOpts, sink chan<- *VnftNFTBurned) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "NFTBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftNFTBurned)
				if err := _Vnft.contract.UnpackLog(event, "NFTBurned", log); err != nil {
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

// ParseNFTBurned is a log parse operation binding the contract event 0x3c176691ca154a2f6fe978a2a633a33ee77dbe2902e67a75400720845a4b2ce1.
//
// Solidity: event NFTBurned(uint256 _tokenId)
func (_Vnft *VnftFilterer) ParseNFTBurned(log types.Log) (*VnftNFTBurned, error) {
	event := new(VnftNFTBurned)
	if err := _Vnft.contract.UnpackLog(event, "NFTBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Vnft contract.
type VnftNFTMintedIterator struct {
	Event *VnftNFTMinted // Event containing the contract specifics and raw log

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
func (it *VnftNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftNFTMinted)
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
		it.Event = new(VnftNFTMinted)
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
func (it *VnftNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftNFTMinted represents a NFTMinted event raised by the Vnft contract.
type VnftNFTMinted struct {
	TokenId               *big.Int
	WithdrawalCredentials []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x996881fd88a98d7719297093e3d0bc1e37ca46ca692ad9dcb15a024eaeba33d8.
//
// Solidity: event NFTMinted(uint256 _tokenId, bytes withdrawalCredentials)
func (_Vnft *VnftFilterer) FilterNFTMinted(opts *bind.FilterOpts) (*VnftNFTMintedIterator, error) {

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return &VnftNFTMintedIterator{contract: _Vnft.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x996881fd88a98d7719297093e3d0bc1e37ca46ca692ad9dcb15a024eaeba33d8.
//
// Solidity: event NFTMinted(uint256 _tokenId, bytes withdrawalCredentials)
func (_Vnft *VnftFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *VnftNFTMinted) (event.Subscription, error) {

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftNFTMinted)
				if err := _Vnft.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x996881fd88a98d7719297093e3d0bc1e37ca46ca692ad9dcb15a024eaeba33d8.
//
// Solidity: event NFTMinted(uint256 _tokenId, bytes withdrawalCredentials)
func (_Vnft *VnftFilterer) ParseNFTMinted(log types.Log) (*VnftNFTMinted, error) {
	event := new(VnftNFTMinted)
	if err := _Vnft.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vnft contract.
type VnftOwnershipTransferredIterator struct {
	Event *VnftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VnftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftOwnershipTransferred)
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
		it.Event = new(VnftOwnershipTransferred)
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
func (it *VnftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftOwnershipTransferred represents a OwnershipTransferred event raised by the Vnft contract.
type VnftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vnft *VnftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VnftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VnftOwnershipTransferredIterator{contract: _Vnft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vnft *VnftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VnftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftOwnershipTransferred)
				if err := _Vnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vnft *VnftFilterer) ParseOwnershipTransferred(log types.Log) (*VnftOwnershipTransferred, error) {
	event := new(VnftOwnershipTransferred)
	if err := _Vnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vnft contract.
type VnftTransferIterator struct {
	Event *VnftTransfer // Event containing the contract specifics and raw log

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
func (it *VnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftTransfer)
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
		it.Event = new(VnftTransfer)
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
func (it *VnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftTransfer represents a Transfer event raised by the Vnft contract.
type VnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Vnft *VnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VnftTransferIterator, error) {

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

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VnftTransferIterator{contract: _Vnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Vnft *VnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftTransfer)
				if err := _Vnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Vnft *VnftFilterer) ParseTransfer(log types.Log) (*VnftTransfer, error) {
	event := new(VnftTransfer)
	if err := _Vnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VnftUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vnft contract.
type VnftUpgradedIterator struct {
	Event *VnftUpgraded // Event containing the contract specifics and raw log

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
func (it *VnftUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VnftUpgraded)
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
		it.Event = new(VnftUpgraded)
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
func (it *VnftUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VnftUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VnftUpgraded represents a Upgraded event raised by the Vnft contract.
type VnftUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vnft *VnftFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VnftUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vnft.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VnftUpgradedIterator{contract: _Vnft.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vnft *VnftFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VnftUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vnft.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VnftUpgraded)
				if err := _Vnft.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vnft *VnftFilterer) ParseUpgraded(log types.Log) (*VnftUpgraded, error) {
	event := new(VnftUpgraded)
	if err := _Vnft.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
