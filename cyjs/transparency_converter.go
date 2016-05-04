package cyjs

import "strconv"

type TransparencyConverter struct {}

const (
	minCy3 = 0.0
	maxCy3 = 255.0
	minCyjs = 0.0
	maxCyjs = 1.0
)

func NewTransparencyConverter() *TransparencyConverter {
	converter := TransparencyConverter{}
	return &converter
}

func (conv TransparencyConverter) Convert(value string) interface{} {

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		// If invalid, simply return fully opaque value
		return maxCyjs
	}

	// Check range
	if number > maxCy3 {
		return maxCyjs
	} else if number < minCy3 {
		return minCyjs
	}

	// Value in valid range
	return number / maxCy3
}
