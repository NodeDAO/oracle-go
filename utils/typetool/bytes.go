// description:
// @author renshiwei
// Date: 2023/6/29

package typetool

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CompareByte32Arrays Compare whether two [][32]bytes are equal
func CompareByte32Arrays(array1, array2 [][32]byte) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array1); i++ {
		if !bytes.Equal(array1[i][:], array2[i][:]) {
			return false
		}
	}

	return true
}

// Byte32ArrToStrArr Convert [][32]byte to []string
func Byte32ArrToStrArr(arr [][32]byte) []string {
	strArr := make([]string, 0, len(arr))
	for _, item := range arr {
		strArr = append(strArr, hexutil.Encode(item[:]))
	}
	return strArr
}
