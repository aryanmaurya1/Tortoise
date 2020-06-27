package main

import "fmt"

func main() {
	var t Table

	t.Header = []string{"A", "B", "C"}
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{3, 4, 5})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	// fmt.Println(EuclideanDistance(t.Rows[0], t.Rows[1]))

	var model KNNModel
	model.Data = t
	fmt.Println(model)
	// fmt.Println(t)
}
