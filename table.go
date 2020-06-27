package main

import "fmt"

// Table : Basic datatype (2D Matrix)
type Table struct {
	Header []string
	Rows   [][]float64
}

// Representation of Table
func (t Table) String() string {
	fmt.Printf("%5s", "SN.")
	for _, value := range t.Header {
		if len(value) > 10 {
			value = value[0:8] + "."
		}
		fmt.Printf("%10s", value)
	}
	fmt.Println()
	var sn int
	for _, value := range t.Rows {
		sn++
		fmt.Printf("%5d", sn)
		for _, cell := range value {
			fmt.Printf("%10.3f", cell)
		}
		fmt.Println()
	}
	return ""
}
