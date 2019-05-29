package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c int
		s       string
	)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}
