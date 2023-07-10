
## 生成abi
```shell
solc --abi --bin <your-contract.sol> -o <output-directory>
```

## 生成合约代码

```shell
abigen --abi=contracts/liq/LiquidStaking.json --pkg=liq --out=contracts/liq/LiquidStaking.go

abigen --abi=contracts/vnft/VNFT.json --pkg=vnft --out=contracts/vnft/vnft.go

abigen --abi=contracts/hashConsensus/MultiHashConsensus.json --pkg=hashConsensus --out=contracts/hashConsensus/multi_hash_consensus.go

abigen --abi=contracts/withdrawOracle/WithdrawOracle.json --pkg=withdrawOracle --out=contracts/withdrawOracle/withdraw_oracle.go

abigen --abi=contracts/operator/operator.json --pkg=operator --out=contracts/operator/operator.go

abigen --abi=contracts/operatorSlash/operatorSlash.json --pkg=operatorSlash --out=contracts/operatorSlash/operator_slash.go

abigen --abi=contracts/withdrawalRequest/withdrawalRequest.json --pkg=withdrawalRequest --out=contracts/withdrawalRequest/withdrawal_request.go

abigen --abi=contracts/largeStaking/LargeStaking.json --pkg=largeStaking --out=contracts/largeStaking/large_staking.go

abigen --abi=contracts/largeStakeOracle/LargeStakeOracle.json --pkg=largeStakeOracle --out=contracts/largeStakeOracle/large_stake_oracle.go
```
