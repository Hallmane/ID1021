package main

import "fmt"

func (m Map) dijkstra(start *City, end *City) (way []string, cost int) {

	pq := newQueue()             // queue for cities to explore
	explored := map[*City]bool{} // cities visited
	trace := make(map[*City]int)

	pq.Set(start, 0) // initial city

	for !pq.Empty() {

		// current next city in queue with its travel time
		cur_city, time_sum := pq.Pop()
		cur_conns := cur_city.connections

		// create path struct from current city and travel time
		path := Path{cur_city, time_sum}

		if cur_city == end { // then back trace path
			cost = path.travel_time
			p_city := path.city

			for p_city != start {
				way = append(way, p_city.name)
				p_city = trace[p_city.name]
				fmt.Println(p_city)
			}
			break
		}

		explored[cur_city] = true

		fmt.Printf("[%d]current node is: %s\n", pq.Len(), cur_city.name)
		for _, conn := range cur_conns {
			// if already visited
			if explored[conn.city] { // or explored[nKey]
				continue
			}

			// if not in queue, add with travel time
			if _, ok := pq.Get(conn.city); !ok {
				trace[conn.city.name] = conn.city // for backtracing
				pq.Set(conn.city, conn.distance)  // add to queue with time
				fmt.Println(conn.city, " added..")
				continue
			}

			fmt.Printf("-> %s ", conn.city.name)
			fmt.Println("old time:", path.travel_time)

			new_time := path.travel_time + conn.distance
			fmt.Println("new time:", new_time)

			if new_time < path.travel_time {
				trace[conn.city.name] = conn.city
				pq.Set(conn.city, new_time)
			}
		}
	}

	way = append(way, start.name)

	for i, j := 0, len(way)-1; i < j; i, j = i+1, j-1 {
		way[i], way[j] = way[j], way[i]
	}

	return way, cost

}
