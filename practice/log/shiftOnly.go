package main

import (
    "fmt"
)

func main() {
    var (
        n     int
        cnt int
    )
    fmt.Scan(&n)
    num := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&num[i])
    }

    flg := true
    for flg {
        for i := 0; i < n; i++ {
            if num[i]%2 == 0 {
                num[i] /= 2
            } else {
                flg = false
            }
        }
        if flg {
            cnt++
        }
    }
    fmt.Println(cnt)
}