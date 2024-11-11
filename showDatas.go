package main

import "fmt"

func showDatas(datas map[string]Measurement, cities *[]string) {
	fmt.Print("{")
	for i, city := range *cities {
		measurement := datas[city]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f",
			city,
			measurement.Min,
			measurement.Sum / float64(measurement.Count),
			measurement.Max,
		)

		if i < len(*cities)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Printf("}\n")
}