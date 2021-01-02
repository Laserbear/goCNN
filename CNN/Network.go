package CNN

import (
	"fmt"
)

type Network struct {
	layers []layer
	loss func(true Matrix, predicted Matrix) float64
	lossPrime func(true Matrix, predicted Matrix) float64
}

type layer interface {
	forwardPass(input Matrix) Matrix
	backPropMatrix (outputError Matrix, alpha float64) Matrix
}

func (n Network) AddLayer(l layer) {
	n.layers = append(n.layers, l)
}

func (n Network) SetLoss(
	loss func(true Matrix, predicted Matrix) float64,lossPrime func(true Matrix, predicted Matrix) float64) {
	n.loss = loss
	n.lossPrime = lossPrime
}

func Mse (true Matrix, predicted Matrix) float64{
	sum := float64(0)
	for i := uint(0); i < true.length; i++ {
		sum += (true.values[i] - predicted.values[i]) * (true.values[i] - predicted.values[i])
	}
	return sum/float64(true.length)
}
func MsePrime (true Matrix, predicted Matrix) float64{
	sum := float64(0)
	for i := uint(0); i < true.length; i++ {
		sum += (true.values[i] - predicted.values[i]) * (true.values[i] - predicted.values[i])
	}
	return sum/float64(true.length)
}


func (n Network) predict (input Matrix) Matrix {
	output := input
	for i:= 0; i < len(n.layers); i++ {
		output = n.layers[i].forwardPass(input)
	}
	return output
}

func (n Network) fit (xTrain []Matrix, yTrain []Matrix, epochs uint, alpha float64) {
	for i := uint(0); i < epochs; i ++ {
		displayError := float64(0)
		for j := range xTrain {
			output := n.predict(xTrain[j])
			displayError += n.loss(yTrain[j], output)

			realError := Matrix{values: []float64{n.lossPrime(yTrain[j], output)}, dimensions: [2]uint{1, 1}, length: 1}

			for i:= len(n.layers) - 1; i >= 0; i-- {
				realError = n.layers[i].backPropMatrix(realError, alpha)
			}


			fmt.Printf("Error is %v", displayError/float64(len(xTrain)))
		}
	}
}

func ReLUFloat(x float64) float64 {
	if x > 0 {
		return x
	} else {
		return 0
	}
}

func ReLUMatrix(m Matrix) Matrix {
	values := make([]float64, m.length)
	for i := uint(0); i < m.length; i++ {
		values[i] = ReLUFloat(m.values[i])
	}
	return Matrix{values: values, dimensions: m.dimensions, length: m.length}
}

