
## 生成合约代码

```shell
abigen --abi=contracts/liq/LiquidStaking.json --pkg=liq --out=contracts/liq/LiquidStaking.go

abigen --abi=contracts/vnft/VNFT.json --pkg=vnft --out=contracts/vnft/vnft.go

abigen --abi=contracts/oracle/HashConsensus.json --pkg=oracle --out=contracts/oracle/hash_consensus.go

abigen --abi=contracts/oracle/WithdrawOracle.json --pkg=oracle --out=contracts/oracle/withdraw_oracle.go
```
