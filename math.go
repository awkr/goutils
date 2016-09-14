package utils

import "math"

func Round(n float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(n*output)) / output
}

func round(n float64) int {
	return int(n + math.Copysign(0.5, n))
}
