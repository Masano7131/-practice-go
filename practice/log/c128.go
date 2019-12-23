package main

import "fmt"

func main() {
	var n, m, k, s, x int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < m; i++ {
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&s)
			s--
			a[s] |= 1 << uint(i)
		}
	}
	p := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&x)
		p |= x << uint(i)
	}
	ans := 0
	for s := 0; s < (1 << uint(n)); s++ {
		t := 0
		for i := 0; i < n; i++ {
			if s>>uint(i)&1 == 1 {
				t ^= a[i]
			}
			if t == p {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
