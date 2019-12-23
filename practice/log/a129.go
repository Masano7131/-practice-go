package main

import "fmt"

func main() {
	var p, q, r int
	fmt.Scan(&p, &q, &r)
	if p+q < p+r {
		if q+r < p+q {
			fmt.Println(q + r)
		} else {
			fmt.Println(p + q)
		}
	} else {
		if q+r < p+r {
			fmt.Println(q + r)
		} else {
			fmt.Println(p + r)
		}
	}
}
