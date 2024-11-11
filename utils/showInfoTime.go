package utils

import (
	"fmt"
	"slices"
	"time"
)

func ShowInfoTime(h []time.Duration, totalTests time.Duration) {
	sumTime := time.Millisecond * 0;
	for i, logTime := range h {
		sumTime += logTime
		fmt.Printf("Time %d -> %v\n", i+1, logTime)
	}

	fmt.Printf("\nMin Time -> %v\n", slices.Min(h))
	fmt.Printf("Max Time -> %v\n\n", slices.Max(h))
	fmt.Printf("Total Time -> %v\n", sumTime)
	fmt.Printf("Average Time -> %v\n", sumTime / totalTests)
}