package main

import (
	"fmt"
	"math"
)

// EuclideanDistance : calculates traditional euclidean distance.
func EuclideanDistance(p1, p2 []float32) float32 {
	// EuclideanDistance = sqrt((p1-q2)^2 + (p2-q2)^2 + ...... + (pn - qn)^2)

	if len(p1) != len(p2) {
		return -1
	}
	var sqrSummation float32
	for i := 0; i < len(p1); i++ {
		sqrSummation += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}
	return float32(math.Sqrt(float64(sqrSummation)))
}

// SquaredMeanError : Calculates the squared mean error of prediction and actual value.
func SquaredMeanError(prediction, actual []float32) float32 {
	/*
		SquaredMeanError = 1/2m * summation((y' - y)^2)
		y' : Prediction
		y  : Actual Value
		m  : Number of datapoints
	*/

	m := len(prediction)
	diff := RowiseSubtract(prediction, actual)
	for i := 0; i < len(diff); i++ {
		diff[i] = diff[i] * diff[i]
	}
	return (1 / (2 * float32(m))) * RowSum(diff)
}

// Gradient : Calculate Gradient based on the predictions made by model.
func Gradient(prediction, actual []float32, t Table) []float32 {
	/*
		dJ/dW = (1 / m) * summation((y' - y) * x)
		y' : Prediction
		y  : Actual Value
		m  : Number of datapoints
	*/

	if !(len(prediction) == len(actual) && len(t.Rows) == len(actual)) {
		fmt.Println("Dimension Mismatch in gradient.")
		return nil
	}
	m := len(t.Rows)
	features := len(t.Rows[0])
	summationRow := make([]float32, features)
	for i := 0; i < m; i++ {
		diff := prediction[i] - actual[i]
		for j := 0; j < features; i++ {
			summationRow[j] = summationRow[j] + (t.Rows[i][j] * diff) // [i][j] = [i][j] + [i][j]*n
		}
	}
	for i := 0; i < features; i++ {
		summationRow[i] = summationRow[i] / float32(m)
	}
	return summationRow
}
