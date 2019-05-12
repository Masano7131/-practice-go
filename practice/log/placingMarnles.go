package main

import(
	"fmt"
)

func main () {
	var (
		input string
		count = 0
	)
	fmt.Scan(&input)
	for _, num := range input {
		if string([]rune {num}) == "1" {
			count ++
		}
	}
	fmt.Println(count)
}