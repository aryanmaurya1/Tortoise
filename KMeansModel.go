package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Point : contains distance from centroid and index of row
type Point struct {
	distance float64
	rowIndex int
}

// KMeansModel : Implements K Nearest Neighbhour clustering algorithm
type KMeansModel struct {
	Name         string
	Data         Table
	Iterations   int // Number of iterations to run
	NumCentroids int
	Centroids    [][]float64
	Clusters     [][]Point
	/*
		Row indexes of table which belongs to a particular cluster.
		Ex :   [[{distance, rowIndex}, {distance, rowIndex}, {distance, rowIndex}], Points for cluster 0
				[i1, i2, i3], Points for cluster 1
				[i1, i2, i3], Points for cluster 2
				.
				.
				.
				[i1, i2, i3]] Indexs for cluster n
	*/
}

func modifiedRow(n float64, row []float64) []float64 {
	// modifies randomly choosen row, using seed n.
	arr := make([]float64, len(row))
	for i := 0; i < len(row); i++ {
		arr[i] = row[i] * 0.5 * n
	}
	return arr
}

func fillCentroids(data *[][]float64, centroids *[][]float64) {
	// fill centroids randomly
	if len(*data) < len(*centroids) {
		*centroids = nil
	}
	rand.Seed(time.Now().Unix())
	var table = make([]int, len(*data), len(*data))
	for i := 0; i < len(*centroids); i++ {
	label:
		randIndex := rand.Int() % len(*data)
		if table[randIndex] == 1 {
			goto label
		}
		table[randIndex] = 1
		(*centroids)[i] = modifiedRow(2, ((*data)[randIndex]))
	}
}

// Init : Initializes the model
func (m *KMeansModel) Init(name string, nCentroids int) {
	m.Name = name
	m.NumCentroids = nCentroids
	m.Centroids = make([][]float64, nCentroids)
	m.Clusters = make([][]Point, nCentroids)
	fillCentroids(&m.Data.Rows, &m.Centroids)
}

func (m *KMeansModel) updateCentroids() {
	var newCentroid = make([][]float64, len(m.Centroids), len(m.Centroids))
	for i := 0; i < len(newCentroid); i++ {
		newCentroid[i] = make([]float64, len(m.Centroids[0]), len(m.Centroids[0]))
		//for j := 0; j < len(m.Clusters); j++ {
		for k := 0; k < len(m.Clusters[i]); k++ {
			rowIndex := (m.Clusters[i][k]).rowIndex
			newCentroid[i] = RowSum(newCentroid[i], m.Data.Rows[rowIndex])
		}
		//}
		for k := 0; k < len(newCentroid[i]); k++ {
			newCentroid[i][k] /= float64(len(m.Clusters[i]))
		}
	}
	m.Centroids = newCentroid
}

func (m *KMeansModel) singlePass() {
	for i := 0; i < len(m.Data.Rows); i++ {
		var minDist float64 = 1000000000.0 // Pseudo +INF
		var centroidNum int = -1
		for j := 0; j < len(m.Centroids); j++ {
			eculidDist := EuclideanDistance(m.Data.Rows[i], m.Centroids[j])
			if eculidDist < minDist {
				minDist = eculidDist
				centroidNum = j
			}
		}
		var p = Point{distance: minDist, rowIndex: i} // storing distance form centroid and row number in table
		m.Clusters[centroidNum] = append(m.Clusters[centroidNum], p)
	}
	// make cluster array empty again for new pass
}

// Run : Runs model for fixed number of iterations
func (m *KMeansModel) Run(iterations int) {
	m.Iterations = iterations
	for i := 0; i < iterations; i++ {
		m.singlePass()
		m.updateCentroids()
		if i != iterations-1 {
			for j := 0; j < len(m.Clusters); j++ {
				m.Clusters[j] = nil // making clusters empty again if not last iteration
			}
		}
	}
}

// Used for printing the model
func (m KMeansModel) String() string {
	fmt.Printf("%20s : %20v \n", "Model Name", m.Name)
	fmt.Printf("%20s : %20v \n", "Iterations", m.Iterations)
	fmt.Printf("%20s : %20v \n", "Number of Centroids", m.NumCentroids)
	fmt.Printf("\n %20s \n", "Attached Table")
	fmt.Println(m.Data)
	fmt.Println(fmt.Sprintf("%20s", "Centroids : "), m.Centroids)
	fmt.Println(fmt.Sprintf("%20s", "Clusters : "), m.Clusters)
	return ""
}
