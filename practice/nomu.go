package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	ft := big.NewInt(1)
	z := big.NewInt(0)
	for i := 0; i < n; i++ {
		t := big.NewInt(0)
		fmt.Scan(&t)
		ft = new(big.Int).Mul(ft, t)
		if !ft.IsInt64() {
			fmt.Println(-1)
			return
		} else if ft == z {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(ft)
}
