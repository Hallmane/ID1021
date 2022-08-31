package main

import (
    "math/rand"
    "fmt"
    "time"
)

func key_value_time(key_len int, value_len int) (float64){
	var (
        //total time.Duration
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
                    break
                }
            }
        }
        t1 := time.Now().Sub(t0)
        t_total += t1
	}
    return float64(t_total.Milliseconds())/float64(sample_size)
}
func main() {
    fmt.Printf("%g\n", key_value_time(10000, 1000))
}

