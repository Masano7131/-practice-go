package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	for i := 0; i < n; i += 2 {
		a += cards[i]
		if i+1 < n {
			b += cards[i+1]
		}
	}
	fmt.Println(a - b)
}
