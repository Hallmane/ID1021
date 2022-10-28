package main

func shoresh(from *City, to *City, max int) *int {
	if max < 0 {
		return nil
	}
	if from == to {
		z := 0
		return &z
	}

	var short *int
	short = &max

	for i := 0; i < len(from.connections); i++ {
		c := from.connections[i]
		distance := shoresh(c.city, to, max-c.distance)

		if distance != nil {
			t := *distance + c.distance
			if t < *short {
				short = &t
			}
		}
	}
	return short
}
