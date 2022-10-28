package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func csv_Parser(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func get_map_from_CSV(filename string) Map {
	map1 := init_map()
	lines, err := csv_Parser(filename)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		c1 := map1.LookUp(line[0])
		c2 := map1.LookUp(line[1])
		dist, _ := strconv.Atoi(line[2])
		c1.Add_Connection(c2, dist)

	}
	return map1
}
