package main

import (
	"fmt"
	"math"
)

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

// RowSum : sums two rows
func RowSum(p1, p2 []float64) []float64 {
	if len(p1) != len(p2) {
		fmt.Println("Row Sum : Length Error")
		return nil
	}
	result := make([]float64, len(p1), len(p1))
	for i := 0; i < len(p1); i++ {
		result[i] = p1[i] + p2[i]
	}
	return result
}
