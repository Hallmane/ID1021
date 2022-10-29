package main

import (
	"fmt"
	"time"
)

func main_test(c1 string, c2 string, max int) {
	map1 := get_map_from_CSV("trains.csv")
	t0 := time.Now()
	dist := shoresh(map1.LookUp(c1), map1.LookUp(c2), max)
	fmt.Printf("[%s -> %s] fastest trip: %d min [%d ms]\n", c1, c2, *dist, time.Since(t0).Milliseconds())
}

func main_test2(c1 string, c2 string, max int) {
	map1 := get_map_from_CSV("trains.csv")
	t0 := time.Now()
	path := new_path()
	dist := path.Shoresh2(map1.LookUp(c1), map1.LookUp(c2), max)
	fmt.Printf("[%s -> %s] fastest trip: %d min [%d µs]\n", c1, c2, *dist, time.Since(t0).Microseconds())
}

func main_test3(c1 string, c2 string, max int) {
	map1 := get_map_from_CSV("trains.csv")
	t0 := time.Now()
	path := new_path()
	dist := path.Shoresh3(map1.LookUp(c1), map1.LookUp(c2), max)
	t1 := time.Since(t0).Microseconds()
	fmt.Printf("[%s -> %s] fastest trip: %d min [%d µs]\n", c1, c2, *dist, t1)
}

func dijkstra_speed(c1, c2 string) {
	map1 := get_map_from_CSV("trains.csv")
	t0 := time.Now()
	start := map1.LookUp(c1)
	end := map1.LookUp(c2)
	map1.dijkstra(start, end)
	t1 := time.Since(t0).Microseconds()
	fmt.Printf("[%s -> %s] fastest trip: [fix] min [%d µs]\n", c1, c2, t1)

}
