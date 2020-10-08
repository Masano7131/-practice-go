package main

import (
	"fmt"
)

func main() {
	var N, M, sum int
	fmt.Scan(&N, &M)
	p := make([]int, N)
	sum = 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		p[i] = a
		sum += a
	}
	var count = 0
	var s = 4 * M
	var lim = sum / s
	for i := 0; i < N; i++ {
		if p[i] > lim {
			count++
		}
		if M > count+N-i-1 {
			fmt.Println("No")
			return
		}
		if count == M {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
	return
}
