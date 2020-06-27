package main

import "math"

// EuclideanDistance : calculates traditional euclidean distance
func EuclideanDistance(p1, p2 []float64) float64 {
	if len(p1) != len(p2) {
		return -1
	}
	var sqrSummation float64
	for i := 0; i < len(p1); i++ {
		sqrSummation += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}
	return math.Sqrt(sqrSummation)
}
