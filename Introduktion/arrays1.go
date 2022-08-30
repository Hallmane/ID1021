package main

import (
	"fmt"
	"time"
	"math/rand"
)

//loops, _ = strconv.ParseInt(os.Args[1], 10, 64)

var (
    slice2 = make([]int, 100000)
)

func array_access_time(loops int, sliceSize int) {
	var (
        total time.Duration
	    mean float64
    )

    slice := make([]int, sliceSize)
    var r4nd = rand.Intn(sliceSize)
    var sum int = 0

	for i := 0; i < loops; i++ {
        t0 := time.Now()
		sum += slice[r4nd]
		total += time.Now().Sub(t0)
	}

	mean = float64(total) / float64(loops)
	fmt.Printf("mean resolution: %g ns\n", mean)
}

func single_read_access_time(n int) (diff time.Duration) {
    var sum int = 0
    var r4nd = rand.Intn(n)

	t0 := time.Now()
	sum += slice2[r4nd]
    return time.Now().Sub(t0)
}

func main() {

	array_access_time(10000, 1000000)

}

//func nanoDiff() {
//	var (total time.Duration
//        loops         = 100
//        sliceSize     = 1000
//	    sum int       = 0
//	    slice1        = make([]int, sliceSize)  
//	    mean float64
//        r4nd = rand.Intn
//    )
//
//    // create slice
//    slice := make([]int, sliceSize)
//	for i := 0; i <= sliceSize; i++ {
//		slice[i] = i
//	}
//
//	for i := 0; i < loops; i++ {
//		n0 := time.Now()
//		sum += slice1[i]
//		total += time.Now().Sub(n0)
//	}
//	mean = float64(total) / float64(loops)
//	fmt.Printf("mean resolution: %g ns\n", mean)
//}
