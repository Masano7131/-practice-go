package main 

import (
	"fmt"
	"strings"
)

func main() {
	var i string
	fmt.Scan(&i)
	fmt.Println(strings.Count(i, "1"))
}