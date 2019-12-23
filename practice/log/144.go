package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	if isPrime(N) {
		fmt.Println(N - 1)
		return
	}
	ina := make([]int, 0)
	primes := make([]int, 0)
	for n := 2; n <= N; n++ {
		flag := true
		for m := 2; m < n; m++ {
			if (n % m) == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, int(n))
		}
	}
	for _, p := range primes {
		if N%p == 0 {
			ina = append(ina, int(p-1+((N/p)-1)))
		}
	}
	fmt.Println(minInt(ina))
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	if x < 2 || x%2 == 0 {
		return false
	}
	for i := 3; i <= x*x; {
		if x%i == 0 {
			return false
		}
		i = i + 2
	}
	return true
}

func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}
