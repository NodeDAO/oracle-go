// desc:
// @author renshiwei
// Date: 2023/4/11 17:49

package abitype

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	StringTy, _       = abi.NewType("string", "", nil)
	BytesTy, _        = abi.NewType("bytes", "", nil)
	Bytes32Ty, _      = abi.NewType("bytes32", "", nil)
	Uint64Ty, _       = abi.NewType("uint64", "", nil)
	Uint96Ty, _       = abi.NewType("uint96", "", nil)
	Uint256Ty, _      = abi.NewType("uint256", "", nil)
	BytesArrayTy, _   = abi.NewType("bytes[]", "bytes[]", nil)
	Bytes32ArrayTy, _ = abi.NewType("bytes32[]", "", nil)
	Uint256ArrayTy, _ = abi.NewType("uint256[]", "", nil)
)
