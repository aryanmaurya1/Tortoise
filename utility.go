package main

import "fmt"

// RowiseSum : sums two rows
func RowiseSum(p1, p2 []float32) []float32 {
	if len(p1) != len(p2) {
		fmt.Println("Row Sum : Length Error")
		return nil
	}
	result := make([]float32, len(p1), len(p1))
	for i := 0; i < len(p1); i++ {
		result[i] = p1[i] + p2[i]
	}
	return result
}

// RowiseSubtract :
func RowiseSubtract(p1, p2 []float32) []float32 {
	if len(p1) != len(p2) {
		fmt.Println("Row Subtract : Length Error")
		return nil
	}
	result := make([]float32, len(p1), len(p1))
	for i := 0; i < len(p1); i++ {
		result[i] = p1[i] - p2[i]
	}
	return result
}

// RowSum :
func RowSum(p []float32) float32 {
	var sum float32
	for i := 0; i < len(p); i++ {
		sum += p[i]
	}
	return sum
}
