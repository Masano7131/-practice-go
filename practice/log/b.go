package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	Fint, _ := strconv.Atoi(S[0:2])
	Sint, _ := strconv.Atoi(S[2:4])
	if Fint == 0 && Sint == 0 {
		fmt.Println("NA")
		return
	}
	if Fint > 12 && Sint > 12 {
		fmt.Println("NA")
		return
	}
	if Fint <= 12 && Sint <= 12 {
		if Fint == 0 {
			fmt.Println("YYMM")
			return
		}
		if Sint == 0 {
			fmt.Println("MMYY")
			return
		}
		fmt.Println("AMBIGUOUS")
	} else if Fint <= 12 && Sint > 12 {
		if Fint == 0 {
			fmt.Println("NA")
			return
		}
		fmt.Println("MMYY")
	} else if Fint > 12 && Sint <= 12 {
		if Sint == 0 {
			fmt.Println("NA")
			return
		}
		fmt.Println("YYMM")
	}
}
