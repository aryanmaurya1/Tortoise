package main

// LinearRegressionModelTrainer : Implementation of multi-variate linear regression
type LinearRegressionModelTrainer struct {
	xData          Table
	yData          []float32
	learningRate   float32
	interations    int
	parameters     []float32
	costs          []float32
	optimizingAlgo func([]float32, []float32, Table) []float32
	costFunc       func([]float32, []float32) float32
}

// LinearRegressionModelPredicter :
type LinearRegressionModelPredicter struct {
	xData       Table
	yData       []float32
	parameters  []float32
	predictions []float32
}
