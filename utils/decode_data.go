package utils

import (
	"strconv"
)

func HexToIntArray(hex []string) ([]int, error) {
	var output []int

	for _, hexByte := range hex {
		n, err := strconv.ParseInt(hexByte, 16, 64)
		if err != nil {
			return output, err
		}
		output = append(output, int(n))
	}

	return output, nil
}
