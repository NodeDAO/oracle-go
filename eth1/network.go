package eth1

var chainIDMap map[string]int

func init() {
	chainIDMap = make(map[string]int)
	chainIDMap[MAINNET] = 1
	chainIDMap[GOERLI] = 5
	chainIDMap[PRATER] = 5
	chainIDMap[SEPOLIA] = 11155111
}

func IsSameNetwork(network string, chainID int) bool {
	return chainIDMap[network] == chainID
}
