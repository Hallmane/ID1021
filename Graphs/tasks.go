package main

import "fmt"

func malmo_Expanding() {
	//main_test2("Malmö", "Lund", 2000)        //
	//main_test2("Malmö", "Helsingborg", 2000) //
	//main_test2("Malmö", "Kalmar", 2000)      //
	//main_test2("Malmö", "Norrköping", 2000)  //
	//main_test2("Malmö", "Södertälje", 2000)  //
	//main_test2("Malmö", "Stockholm", 2000)   //
	//main_test2("Malmö", "Gävle", 2000)       //
	//main_test2("Malmö", "Sundsvall", 2000)   //
	//main_test2("Malmö", "Östersund", 2000) //
	//main_test2("Malmö", "Gällivare", 2000) //
	//main_test2("Malmö", "Kiruna", 2000)    //
	main_test2("Malmö", "Borlänge", 2000) //
}

func test_csv_parser() {
	map1 := get_map_from_CSV("trains.csv")
	map1.String()
}

func test_city_string() {
	map1 := get_map_from_CSV("trains.csv")
	for _, c := range map1.cities {
		c.String()
	}
}

func test_Get_Connections() {
	map1 := get_map_from_CSV("trains.csv")
	city := map1.cities[7]
	conns := city.Get_Connections()
	conns[0].String()
}

func test_HashMe() {
	fmt.Println(HashMe("Stockholm"))
}

func test_LookUp() {
	map1 := get_map_from_CSV("trains.csv")
	fmt.Println(map1)
	map1.cities[4].String()

}

func test_shoresh() {
	map1 := get_map_from_CSV("trains.csv")

	c1 := map1.LookUp("Gävle")
	c2 := map1.LookUp("Östersund")

	time := shoresh(c1, c2, 400)

	fmt.Println(*time)
}

func shoresh3_test() {
	main_test3("Malmö", "Kiruna", 10000)
}

/*
			DATA DUMP
[Malmö -> Lund] fastest trip: 13 min [0 ms]
[Malmö -> Helsingborg] fastest trip: 48 min [372 ms]
[Malmö -> Kalmar] fastest trip: 160 min [323 ms]
[Malmö -> Norrköping] fastest trip: 193 min [164 ms]
[Malmö -> Södertälje] fastest trip: 252 min [118 ms]
[Malmö -> Stockholm] fastest trip: 273 min [137 ms]
[Malmö -> Gävle] fastest trip: 383 min [130 ms]
[Malmö -> Borlänge] fastest trip: 410 min [132743 µs]
[Malmö -> Sundsvall] fastest trip: 600 min [371 ms]
[Malmö -> Östersund] fastest trip: 612 min [355513 µs]
[Malmö -> Gällivare] fastest trip: 1095 min [520978 µs]
[Malmö -> Kiruna] fastest trip: 1162 min [518339 µs]

*/
