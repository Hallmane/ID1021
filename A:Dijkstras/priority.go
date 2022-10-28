package main

type PrQueue struct {
	keys  []*City
	nodes map[*City]int
}

// city has adjacency built in thanks to []connections
type Path struct {
	city        *City // name, []connections
	travel_time int   //sum of weights
}

func (pq PrQueue) Len() int                { return len(pq.keys) }
func (pq PrQueue) Empty() bool             { return len(pq.keys) == 0 }
func (pq PrQueue) Less(p_A, p_B Path) bool { return p_A.travel_time < p_B.travel_time }

func (pq *PrQueue) Set(city *City, length int) {
	if _, ok := pq.nodes[city]; !ok {
		pq.keys = append(pq.keys, city)
	}
	pq.nodes[city] = length
}
func (pq *PrQueue) Get(key *City) (prio int, ok bool) {
	prio, ok = pq.nodes[key]
	return
}

// pops next item and updates queue
func (pq *PrQueue) Pop() (key *City, prio int) {
	key, keys := pq.keys[0], pq.keys[1:]
	pq.keys = keys

	prio = pq.nodes[key]

	delete(pq.nodes, key)
	return key, prio

}

func makePath(city *City, time int) *Path {
	return &Path{city: city, travel_time: time}
}

func newQueue() *PrQueue {
	var pq PrQueue
	pq.nodes = make(map[*City]int)
	return &pq
}
