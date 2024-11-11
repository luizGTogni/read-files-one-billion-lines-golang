package utils

import (
	"bufio"
	"os"
)

func ScannerData[T any](
	data *os.File, 
	datas map[string]T, 
	create func(s *bufio.Scanner, d map[string]T),
)  {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		create(scanner, datas);
	}

}