package main

import (
	"fmt"
	"time"
	"math/rand"
)
func firstTask() {
    for i:=0; i<10; i++ {
        n0 := time.Now()
        n1 := time.Now().Sub(n0)
        fmt.Printf("Resolution:  %d nanoseconds\n", n1)
    }
}

func array_access_time(loops int, n int) {
	var (
        //total time.Duration
        total_time_1 time.Duration
        total_time_2 time.Duration
	    mean float64
        indx  = make([]int, n)
        slice = make([]int, n)
    )


    for i:=0; i<n; i++ {
        var r4nd = rand.Intn(n)
        slice[i] = 1
        indx[i] = r4nd
    }

    var sum int = 0
    //var dummy_add int = 0

    t_start := time.Now()
	for i := 0; i < loops; i++ {
	    for j := 0; j < n; j++ {
		    sum += slice[indx[j]]
        }
	}
    diff_1 := time.Now().Sub(t_start)

	total_time_1 += diff_1

    t_start = time.Now()
	for i := 0; i < loops; i++ {
	    for j := 0; j < n; j++ {
		    sum += 1
        }
	}
    diff_2 := time.Now().Sub(t_start)
	total_time_2 += diff_2
    
	mean = float64(total_time_1 - total_time_2) / float64(n*loops)
	fmt.Printf("Access time: %g ns\n", mean)
    //median := slice[n/2]
	//fmt.Printf("median resolution: %g ns\n", median)
}
func main() {
    var loops = 100
    var smallSize  = 10000
    var mediumSize = 100000
    var largeSize  = 1000000

    fmt.Printf(" array size: %d\n loops: %d\n", smallSize, loops)
	array_access_time(loops, smallSize)
    fmt.Printf(" array size: %d\n loops: %d\n", mediumSize, loops)
	array_access_time(loops, mediumSize)
    fmt.Printf(" array size: %d\n loops: %d\n", largeSize, loops)
	array_access_time(loops, largeSize)

    firstTask()
}

