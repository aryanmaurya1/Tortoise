package main

import (
	"math"
)

// EuclideanDistance : calculates traditional euclidean distance
func EuclideanDistance(p1, p2 []float32) float32 {
	if len(p1) != len(p2) {
		return -1
	}
	var sqrSummation float32
	for i := 0; i < len(p1); i++ {
		sqrSummation += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}
	return float32(math.Sqrt(float64(sqrSummation)))
}

// SquaredMeanError :
func SquaredMeanError(prediction, actual []float32) float32 {
	m := len(prediction)
	diff := RowiseSubtract(prediction, actual)
	for i := 0; i < len(diff); i++ {
		diff[i] = diff[i] * diff[i]
	}
	return (1 / (2 * float32(m))) * RowSum(diff)
}

// Gradient :
func Gradient(prediction, actual []float32, t Table) []float32 {
	return nil

}
