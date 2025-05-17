package conversions

import (
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		p, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return prices, err
		}
		prices[lineIndex] = p
	}
	return prices, nil
}
