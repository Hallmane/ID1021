package main

import (
    "fmt"
    "time"
)
var (
    //n0 float64 = float64(time.Now().UnixNano())
    //n1 float64 = float64(time.Now().UnixNano())
    //n0, n1 float64 = float64(time.Now().UnixNano()), float64(time.Now().UnixNano())
    given = [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    sum int = 0;
)

func main() {
    for i := 0; i < 10; i++ {
        n0 := float64(time.Now().UnixNano())
        sum += given[i]
        n1 := float64(time.Now().UnixNano())
        fmt.Printf(" resolution  %g   nanoseconds\n", (n1- n0))
    }
}
