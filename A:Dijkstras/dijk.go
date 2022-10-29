package main

type Path struct {
	city        *City // name, []connections
	travel_time int   //sum of weights
}

// aint workin
func (m Map) dijkstra(start, end *City) (way []string, path_time int) {

	pq := newQueue()                // queue for cities to explore
	explored := map[*City]bool{}    // cities visited
	previous := make(map[int]*City) // mapping of int->c
	c_index := make(map[*City]int)  // mapping of c->int

	pq.Set(start, 0) // initial city

	for !pq.Empty() {

		// current next city in queue with its travel time
		cur_city, time_sum := pq.Pop()

		// create path struct from current city and travel time
		path := Path{cur_city, time_sum}

		if cur_city == end { // then back trace path
			path_time = path.travel_time
			//p_city := path.city

			//for p_city != start {
			//	//fmt.Println(way, "<way, path.city>", p_city)
			//	way = append(way, p_city.name)
			//	p_city = previous[c_index[cur_city]]
			//	fmt.Println(len(previous), "previous   city")
			//}

			break
		}

		explored[cur_city] = true

		for _, conn := range cur_city.connections {
			// if already visited
			if explored[conn.city] { // or explored[nKey]
				continue
			}

			// if not in queue, add with travel time
			if _, ok := pq.Get(conn.city); !ok {
				previous[c_index[conn.city]] = conn.city          // for backtracing
				pq.Set(conn.city, path.travel_time+conn.distance) // add to queue with time
				continue
			}

			queue_time, _ := pq.Get(conn.city)
			new_time := path.travel_time + conn.distance

			if new_time < queue_time {
				///if new_time < path.travel_time {
				previous[c_index[conn.city]] = conn.city
				pq.Set(conn.city, new_time)
			}
		}
	}

	way = append(way, start.name)

	for i, j := 0, len(way)-1; i < j; i, j = i+1, j-1 {
		way[i], way[j] = way[j], way[i]
	}

	return

}
