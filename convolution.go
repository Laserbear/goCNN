package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	defer timeMeasurement(time.Now())

	rows, err := strconv.Atoi(os.Args[1])
	if err != nil {}
	//fmt.Printf("Rows: %d \n", rows)
	cols, err := strconv.Atoi(os.Args[2])
	if err != nil {}
	//fmt.Printf("Cols: %d \n", cols)
	var K = []int{1, 0 , -1}  //reverse filter from [-1, 0, 1]
	matrix := [10][10]uint8{}
	Dy := [10][10]int{}
	Dx := [10][10]int{}
	for i:= 0; i < rows; i++  {
		for j := 0; j < cols; j++ {
			matrix[i][j] = uint8(rand.Intn(255))
			fmt.Print(matrix[i][j])
			fmt.Printf(" ")
		}
	}
	//fmt.Printf("%v", matrix[:rows][:cols])

	var slice = [3]uint8{}
	for i:= 0; i < rows; i++  {
		for j := 0; j < cols - 3; j++ {
			for l := 0; l < 3; l++ {
				slice[l] = matrix[i][j+l]
			}
			Dx[i][j] = dotproduct(slice, K)
		}
	}

	for i := 0; i < rows - 3; i++ {
		for j := 0; j < cols; j++ {
			for l := 0; l < 3; l++ {
				slice[l] = matrix[i+l][j]
			}
			Dy[j][i] = dotproduct(slice, K)

		}
	}

	// This is so wrong lmao Dx and Dy should be 2D
	fmt.Printf("Dx%v\n", Dx[:rows])
	fmt.Printf("Dy%v\n", Dy[:cols])


}
func dotproduct(A [3]uint8, B []int) (dotproduct int) {
	sum := 0
	for i := range A {
		sum += int(A[i]) * B[i]
	}
	return sum
}

func timeMeasurement(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}


