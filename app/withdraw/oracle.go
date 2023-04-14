// desc:
// @author renshiwei
// Date: 2023/4/11 15:12

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func EncodeReportData(reportData *withdrawOracle.WithdrawOracleReportData) ([32]byte, error) {
	// Code the data using the Pack method of the abi.Arguments structure
	encodedData, err := reportDataArguments.Pack(reportData)
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
