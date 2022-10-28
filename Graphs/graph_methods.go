package main

import "fmt"

func init_map() Map {
	return Map{}
}

func init_city() City {
	return City{connections: make([]Connection, 0)}
}

func (origin *City) Add_Connection(dest *City, dist int) {
	connection_origin := Connection{dest, dist}
	connection_dest := Connection{origin, dist}

	origin.connections = append(origin.connections, connection_origin)
	dest.connections = append(dest.connections, connection_dest)
}

// adds city to next spot in hashed array that is nil
func (m *Map) Add_To_Map(c *City) {
	step := 0
	for m.cities[HashMe(c.name)+step] != nil {
		step++
	}
	m.cities[HashMe(c.name)+step] = c
}

// traverses the hashed index of the hashed array until the
// specified name is found
func (m *Map) LookUp(searched_name string) *City {
	city_hash := HashMe(searched_name)

	for step := 0; m.cities[city_hash+step] != nil && city_hash+step < len(m.cities); step++ {
		if m.cities[city_hash+step] == nil {
			break
		}

		if m.cities[city_hash+step].name == searched_name {
			return m.cities[city_hash+step]
		}
	}

	new_city := init_city()
	new_city.name = searched_name
	m.Add_To_Map(&new_city)
	return &new_city
}

func (city City) Get_Connections() []*City {
	connections := make([]*City, 0)
	for _, conns := range city.connections {
		connections = append(connections, conns.city)
	}
	return connections
}

func (m Map) String() {
	for _, c := range m.cities {
		fmt.Printf("city: %v\n", c.name)
	}
}

func (c City) String() {
	fmt.Printf("%s\n", c.name)
	for _, conn := range c.connections {
		fmt.Printf("-> %s [%d mins]\n", conn.city.name, conn.distance)
	}
}
