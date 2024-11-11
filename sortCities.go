package main

import "sort"

func sortCities(datas map[string]Measurement, cities *[]string) {
	for city := range datas {
		*cities = append(*cities, city)
	}

	sort.Strings(*cities)
}