package tortoise

// LinearRegressionModelTrainer : Implementation of multi-variate linear regression.
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

// LinearRegressionModelPredicter : Used for making prediction on trained model.
// Parameters are transferred from a trained model (structure) by calling TransferParams functions.
type LinearRegressionModelPredicter struct {
	xData       Table
	parameters  []float32
	predictions []float32
}
