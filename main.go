package main

import (
	"project-advanced-concepts/errors"
	"project-advanced-concepts/utils"
	"sync"
	"time"
)

const AMOUNT_TESTS_PERFOMANCE = 1

func measurementsTempByCities(verbose bool) {
	measurements, _, err := utils.OpenFile("datas/measurements.txt")
	errors.HandlerError(err)
	defer measurements.Close()

	datas := make(map[string]Measurement, 1000000)

	utils.ScannerData(measurements, datas, createMeasurement)

	cities := make([]string, 0, len(datas))
	sortCities(datas, &cities)

	if verbose {
		showDatas(datas, &cities)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(AMOUNT_TESTS_PERFOMANCE)

	history := make([]time.Duration, 0, AMOUNT_TESTS_PERFOMANCE)
	for range AMOUNT_TESTS_PERFOMANCE {
		go func() {
			start := time.Now()
			defer wg.Done()
			measurementsTempByCities(false)
			history = append(history, time.Since(start))
		}()
	}

	wg.Wait()
	utils.ShowInfoTime(history, AMOUNT_TESTS_PERFOMANCE)	
}