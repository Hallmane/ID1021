package main

func HashMe(cityName string) int {
	h := 7
	mod := 541
	for i := 1; i < len(cityName); i++ {
		h = (h * 31 % mod) + int(cityName[i])
	}
	return h % mod
}
