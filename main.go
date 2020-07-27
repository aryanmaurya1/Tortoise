package main

import (
	"fmt"
	"tortoise/tortoise"
)

func main() {
	var t tortoise.Table

	t.Header = []string{"A", "B", "C"}
	t.Rows = append(t.Rows, []float32{1, 2, 3})
	t.Rows = append(t.Rows, []float32{3, 4, 5})
	t.Rows = append(t.Rows, []float32{6, 7, 8})
	t.Rows = append(t.Rows, []float32{17, 72, 78})
	t.Rows = append(t.Rows, []float32{1, 2, 3})
	t.Rows = append(t.Rows, []float32{1, 2, 3})
	t.Rows = append(t.Rows, []float32{1, 2, 3})
	t.Rows = append(t.Rows, []float32{100, 200, 300})
	t.Rows = append(t.Rows, []float32{0, 0, 0})

	var model tortoise.KMeansModel
	model.Data = t
	model.Init("Dummy Model", 3)
	model.Run(100000)
	fmt.Println(model)
}
