
## 生成abi
```shell
solc --abi --bin <your-contract.sol> -o <output-directory>
```

## 生成合约代码

```shell
abigen --abi=contracts/liq/LiquidStaking.json --pkg=liq --out=contracts/liq/LiquidStaking.go

abigen --abi=contracts/vnft/VNFT.json --pkg=vnft --out=contracts/vnft/vnft.go

abigen --abi=contracts/hashConsensus/HashConsensus.json --pkg=hashConsensus --out=contracts/hashConsensus/hash_consensus.go

abigen --abi=contracts/withdrawOracle/WithdrawOracle.json --pkg=withdrawOracle --out=contracts/withdrawOracle/withdraw_oracle.go
```