package main

import (
	"fmt"
)

func main() {
	var (
		a int
	)
	fmt.Scan(&a)
	H := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&H[i])
	}
	j := 0
	k := 0
	for i := 0; i < a-1; i++ {
		if H[i] >= H[i+1] {
			j++
		} else {
			if k <= j {
				k = j
				j = 0
			}
			j = 0
		}
	}
	if k < j {
		fmt.Println(j)
	} else {
		fmt.Println(k)
	}
}
