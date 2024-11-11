package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Measurement struct {
	Min 	float64
	Max 	float64
	Sum 	float64
	Count int
}


func createMeasurement(scanner *bufio.Scanner, datas map[string]Measurement) {
	rawData := scanner.Text()
	splitData := strings.Split(rawData, ";");
	city := splitData[0]
	rawTemp := splitData[1]

	temp, _ := strconv.ParseFloat(rawTemp, 64)

	measurement, ok := datas[city]
	if !ok {
		measurement = Measurement{
			Min: temp,
			Max: temp,
			Sum: temp,
			Count: 1,
		}
	} else {
		measurement.Min = min(measurement.Min, temp)
		measurement.Max = max(measurement.Max, temp)
		measurement.Sum += temp
		measurement.Count++
	}

	datas[city] = measurement
}