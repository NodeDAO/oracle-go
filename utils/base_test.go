// desc:
// @author renshiwei
// Date: 2023/4/10 18:32

package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

func TestMakeArr(t *testing.T) {
	arr := make([]int, 1)
	arr[0] = 1
	fmt.Printf("arr:%+v\n", arr)
	fmt.Printf("arr len:%+v\n", len(arr))
	fmt.Printf("arr cap:%+v\n", cap(arr))

	arr2 := make([]int, 1, 1)
	arr2[0] = 1
	fmt.Printf("arr2:%+v\n", arr2)
	fmt.Printf("arr2 len:%+v\n", len(arr2))
	fmt.Printf("arr2 cap:%+v\n", cap(arr2))

	arr3 := make([]int, 1, 2)
	arr3[0] = 1
	//arr3[0] = 1  err
	arr3 = append(arr3, 1)
	fmt.Printf("arr3:%+v\n", arr3)
	fmt.Printf("arr3 len:%+v\n", len(arr3))
	fmt.Printf("arr3 cap:%+v\n", cap(arr3))
}

func TestDecimalForBigInt(t *testing.T) {
	fmt.Println(decimal.NewFromBigInt(big.NewInt(111), 0).String())
	fmt.Println(decimal.NewFromBigInt(big.NewInt(111), 1).String())
}
