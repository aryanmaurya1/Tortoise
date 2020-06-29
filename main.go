package main

import (
	"fmt"
)

func main() {
	var t Table

	t.Header = []string{"A", "B", "C"}
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{3, 4, 5})
	t.Rows = append(t.Rows, []float64{6, 7, 8})
	t.Rows = append(t.Rows, []float64{17, 72, 78})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{1, 2, 3})
	t.Rows = append(t.Rows, []float64{100, 200, 300})
	t.Rows = append(t.Rows, []float64{0, 0, 0})
	// fmt.Println(EuclideanDistance(t.Rows[0], t.Rows[1]))

	var model KMeansModel
	model.Data = t
	model.Init("Dummy Model", 3)
	// fmt.Println(model)
	model.Run(10000000)
	fmt.Println(model)
	// fmt.Println(t)
}
