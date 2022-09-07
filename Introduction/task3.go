package main

import (
    "math/rand"
    "fmt"
    "time"
)

func key_value_time(key_len int, value_len int) (float64, int){
    rand.Seed(time.Now().UTC().UnixNano())
	var (
        keys = make([]int, key_len)
        values = make([]int, value_len)
        sum int = 0
        sample_size int = 100;
        t_total time.Duration
    )

	for i := 0; i < sample_size; i++ {
        for j:=0; j<key_len; j++ {
            keys[j] = rand.Intn(10*value_len)
        }
        for j:=0; j<value_len; j++ {
            values[j] = rand.Intn(10*value_len)
        }
        t0 := time.Now()
	    for k := 0; k < key_len; k++ {
            key := keys[k]
            for v := 0; v < value_len; v++ {
                if values[v] == key {
                    sum++
                }
            }
        }
        t1 := time.Now().Sub(t0)
        t_total += t1
	}
    //fmt.Printf("%d\n",sum)
    return float64(t_total)/float64(sample_size), sum
}
func main() {
    var key_len, value_len int = 100, 100
    time_start := time.Now()

    for x:=10; x<90000000; x+=10 {
        timeS, _:= key_value_time(key_len, value_len)
        key_len, value_len = key_len + x, value_len + x
        //fmt.Printf("%g nanoseconds\n", timeNS)
        //fmt.Printf("%d sum\n", sum)
        //fmt.Printf("%d %g\n", sum, time)
        //fmt.Printf("%d %g\n", value_len, time)
        fmt.Printf("%d %f\n", value_len, timeS)

        if time.Since(time_start).Seconds() >= 60 {
            break
        }
    }
}

















