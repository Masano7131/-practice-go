package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mochis := make([]int, n)
	overlapm := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&mochis[i])
	}
	m := make(map[int]bool)
	for _, mochi := range mochis {
		if !m[mochi] {
			m[mochi] = true
			overlapm = append(overlapm, mochi)
		}
	}
	fmt.Println(len(overlapm))
}
