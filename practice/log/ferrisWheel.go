package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A >= 13 {
		fmt.Println(B)
	} else if A <= 5 {
		fmt.Println(0)
	} else {
		fmt.Println(B / 2)
	}

}
