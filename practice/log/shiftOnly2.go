package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readInt(s *bufio.Scanner) int {
	if !s.Scan() {
		return 0
	}
	n, _ := strconv.Atoi(s.Text())
	return n
}

func checkEven(a int) int {
	as := 0
	for i := 2; i <= a; i *= 2 {
		if a%i != 0 {
			return as
		}
		as++
	}
	return as
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := readInt(sc)
	m := 0
	for i := 0; i < n; i++ {
		mm := checkEven(readInt(sc))
		if i == 0 {
			m = mm
			continue
		}
		m = min(m, mm)
	}
	fmt.Println(m)
}
