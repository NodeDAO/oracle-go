// desc:
// @author renshiwei
// Date: 2023/4/11 15:12

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

// REPORT_DATA_TYPE  (uint256,uint256,uint256,uint256,uint256,uint256,(uint64,uint96,uint96)[],(uint64,uint96,uint96)[],uint256[],uint256[]) data
var (
	withdrawInfosArgumentMarshaling      = []abi.ArgumentMarshaling{{Type: "uint64"}, {Type: "uint96"}, {Type: "uint96"}}
	exitValidatorInfosArgumentMarshaling = []abi.ArgumentMarshaling{{Type: "uint64"}, {Type: "uint96"}, {Type: "uint96"}}

	//withdrawInfosType, _      = abi.NewType("tuple", "", withdrawInfosArgumentMarshaling)
	//exitValidatorInfosType, _ = abi.NewType("tuple", "", exitValidatorInfosArgumentMarshaling)

	reportDataTypeArgumentMarshaling = []abi.ArgumentMarshaling{
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "uint256"},
		{Type: "tuple", Components: withdrawInfosArgumentMarshaling},
		{Type: "tuple", Components: exitValidatorInfosArgumentMarshaling},
		{Type: "uint256[]"},
		{Type: "uint256[]"},
	}

	reportDataType, _   = abi.NewType("tuple", "", reportDataTypeArgumentMarshaling)
	reportDataArguments = abi.Arguments{{Type: reportDataType}}
)

func EncodeReportData(reportData *withdrawOracle.WithdrawOracleReportData) ([32]byte, error) {
	json, err := abi.JSON(strings.NewReader("[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clVaultBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clSettleAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportExitedCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"clReward\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"clCapital\",\"type\":\"uint96\"}],\"internalType\":\"struct WithdrawInfo[]\",\"name\":\"withdrawInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"exitTokenId\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"exitBlockNumber\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"slashAmount\",\"type\":\"uint96\"}],\"internalType\":\"struct ExitValidatorInfo[]\",\"name\":\"exitValidatorInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"delayedExitTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"largeExitDelayedRequestIds\",\"type\":\"uint256[]\"}],\"internalType\":\"struct WithdrawOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"))
	encodedData, err := json.Methods["submitReportData"].Inputs.Pack(reportData)

	// Code the data using the Pack method of the abi.Arguments structure
	//encodedData, err := reportDataArguments.Pack(reportData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode report data err.")
	}

	// Calculate the keccak 256 hash value
	hash := crypto.Keccak256(encodedData)
	var res [32]byte
	copy(res[:], hash)

	return res, nil
}

func (v *Oracle) Paused(ctx context.Context) (bool, error) {
	paused, err := contracts.WithdrawOracleContract.Contract.Paused(nil)
	if err != nil {
		return true, errors.Wrap(err, "Failed to get WithdrawOracleContract Paused.")
	}

	return paused, nil
}

func (v *Oracle) GetConsensusContract(ctx context.Context) (*hashConsensus.HashConsensus, error) {
	address, err := v.GetConsensusContractAddress(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	contract, err := hashConsensus.NewHashConsensus(address, eth1.ElClient.Client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get withdrawOracleHelper HashConsensus.")
	}

	return contract, nil
}

func (v *Oracle) GetConsensusContractAddress(ctx context.Context) (common.Address, error) {
	addr, err := contracts.WithdrawOracleContract.Contract.GetConsensusContract(nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to get WithdrawOracleContract GetConsensusContract.")
	}
	return addr, nil
}

func (v *Oracle) CheckContractVersions(ctx context.Context) error {
	oracleContractVersion, err := contracts.WithdrawOracleContract.Contract.GetContractVersion(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract GetContractVersion.")
	}

	consensusVersion, err := contracts.WithdrawOracleContract.Contract.GetConsensusVersion(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract GetConsensusVersion.")
	}

	if oracleContractVersion.Int64() != CONTRACT_VERSION || consensusVersion.Int64() != CONSENSUS_VERSION {
		return errors.Errorf(`Incompatible Oracle version err. 
			Expected contract version %v got %v.
			Expected consensus version %v got %v.
		`, oracleContractVersion.Int64(), CONTRACT_VERSION, consensusVersion.Int64(), CONSENSUS_VERSION)
	}

	return nil
}

func (v *Oracle) IsContractReportable(ctx context.Context) (bool, error) {
	isReport, err := v.IsMainDataSubmitted(ctx)
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	return !isReport, nil
}

func (v *Oracle) IsMainDataSubmitted(ctx context.Context) (bool, error) {
	processingState, err := v.GetProcessingState(ctx)
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	return processingState.DataSubmitted, nil
}

func (v *Oracle) GetProcessingState(ctx context.Context) (*withdrawOracle.WithdrawOracleProcessingState, error) {
	processingState, err := contracts.WithdrawOracleContract.Contract.GetProcessingState(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get WithdrawOracleContract GetProcessingState.")
	}
	return &processingState, nil
}

func (v *Oracle) simulatedSubmitReportData(ctx context.Context, reportData withdrawOracle.WithdrawOracleReportData, consensusVersion *big.Int) error {
	chainID, err := eth1.ElClient.Client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get chainID.")
	}

	balance := big.NewInt(5e18)

	simulatedClient, err := eth1.NewSimulatedClient(chainID, balance, 2000000)
	if err != nil {
		return errors.Wrap(err, "Failed to get simulatedClient.")
	}

	opts, err := eth1.KeyTransactOpts(chainID)
	if err != nil {
		return errors.Wrapf(err, "")
	}

	address := opts.From

	withdrawOracleTransactor, err := withdrawOracle.NewWithdrawOracleTransactor(address, simulatedClient)
	if err != nil {
		return errors.Wrapf(err, "Failed to get withdrawOracleTransactor.")
	}

	tx, err := withdrawOracleTransactor.SubmitReportData(opts, reportData, consensusVersion)
	if err != nil {
		return errors.Wrap(err, "simulated WithdrawOracle SubmitReportData err.")
	}

	// Wait for the transaction to complete
	if _, err = bind.WaitMined(context.Background(), simulatedClient, tx); err != nil {
		return errors.Wrapf(err, "Failed to simulated WaitMined submit report data. tx hash:%s", tx.Hash().String())
	}
	logger.Info("simulated tx Send report data success.",
		zap.String("tx hash", tx.Hash().String()),
		zap.String("consensusVersion", consensusVersion.String()),
	)

	return nil
}
