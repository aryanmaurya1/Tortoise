package main

import "fmt"

// KNNModel : Implements K Nearest Neighbhour clustering algorithm
type KNNModel struct {
	Name            string
	Data            Table
	Iterations      int // Number of iterations to run
	NumCentroids    int
	RandomCentroids bool
	Centroids       []float64
	Clusters        [][]float64
	/*
		Row indexes of table which belongs to a particular cluster.
		Ex :   [[i1, i2, i3], Indexs for cluster 1
				[i1, i2, i3], Indexs for cluster 2
				[i1, i2, i3], Indexs for cluster 3
				.
				.
				.
				[i1, i2, i3]] Indexs for cluster n
	*/
}

// Used for printing the model
func (m KNNModel) String() string {
	fmt.Printf("%20s : %20v \n", "Model Name", m.Name)
	fmt.Printf("%20s : %20v \n", "Iterations", m.Iterations)
	fmt.Printf("%20s : %20v \n", "Number of Centroids", m.NumCentroids)
	fmt.Printf("%20s : %20v \n", "Random Centroids", m.RandomCentroids)
	fmt.Printf("\n %20s \n", "Attached Table")
	fmt.Println(m.Data)
	fmt.Println(fmt.Sprintf("%20s", "Centroids : "), m.Centroids)
	fmt.Println(fmt.Sprintf("%20s", "Clusters : "), m.Clusters)
	return ""
}
