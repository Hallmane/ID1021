package main

const MAX_PATH_LENGTH = 54

type Paths struct {
	cities [MAX_PATH_LENGTH]*City
	sp     int
}

func new_path() *Paths {
	return &Paths{sp: 0}
}

func (path *Paths) Shoresh2(from *City, to *City, max int) *int {
	if max < 0 {
		return nil
	}
	if from == to {
		z := 0
		return &z
	}
	for i := 0; i < path.sp; i++ {
		if path.cities[i] == from {
			return nil
		}
	}

	path.cities[path.sp] = from
	path.sp++

	var short *int
	short = &max

	for i := 0; i < len(from.connections); i++ {
		c := from.connections[i]
		distance := path.Shoresh2(c.city, to, max-c.distance)

		if distance != nil {
			t := *distance + c.distance
			//fmt.Println(t, "< path | connection > ", c)

			if t < *short {
				short = &t
			}
		}
	}
	path.cities[path.sp] = nil
	path.sp--
	return short
}

func (path *Paths) Shoresh3(from *City, to *City, max int) *int {
	if max < 0 {
		return nil
	}
	if from == to {
		z := 0
		return &z
	}
	for i := 0; i < path.sp; i++ {
		if path.cities[i] == from {
			return nil
		}
	}

	path.cities[path.sp] = from
	path.sp++

	var short *int
	short = &max

	for i := 0; i < len(from.connections); i++ {
		c := from.connections[i]
		distance := path.Shoresh3(c.city, to, max-c.distance)

		if distance != nil {
			t := *distance + c.distance

			if t < *short {
				short = &t
				max = *short + c.distance
			}
		}
	}
	path.cities[path.sp] = nil
	path.sp--
	return short
}
