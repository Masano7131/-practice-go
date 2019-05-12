package main

import "fmt"

func sumi(n int) (s int) {
	s = 0
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return s
}
func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	sum := 0
	for i := 1; i <= n; i++ {
		if a <= sumi(i) && sumi(i) <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
