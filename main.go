package main


func main() {
	//disgusting but should learn XOR
	//TODO: make matrix initialization cleaner; maybe just pass a list and dimensions in?
	Matrix00 := Matrix{values: []float64{0, 0}, dimensions: [2]uint{2, 1}, length: 2}
	Matrix01 := Matrix{values: []float64{0, 1}, dimensions: [2]uint{2, 1}, length: 2}
	Matrix10 := Matrix{values: []float64{1, 0}, dimensions: [2]uint{2, 1}, length: 2}
	Matrix11 := Matrix{values: []float64{1, 1}, dimensions: [2]uint{2, 1}, length: 2}
	label00 := Matrix{values: []float64{0}, dimensions: [2]uint{1, 1}, length: 1}
	label01 := Matrix{values: []float64{1}, dimensions: [2]uint{1, 1}, length: 1}
	label10 := Matrix{values: []float64{1}, dimensions: [2]uint{1, 1}, length: 1}
	label11 := Matrix{values: []float64{0}, dimensions: [2]uint{1, 1}, length: 1}
	x_data := []Matrix{Matrix00, Matrix01, Matrix10, Matrix11}
	y_labels := []Matrix{label00, label01, label10, label11}
	net := Network{layers: []layer{}, loss: Mse, lossPrime: MsePrime}
	net.AddLayer(FullyConnectedLayer{weights: randomFloatMatrix(2, 3), bias: randomFloatMatrix(1, 3)})
	net.AddLayer(ActivationLayer{activation: ReLUMatrix, activationPrime: ReLUMatrix})
	net.AddLayer(FullyConnectedLayer{weights: randomFloatMatrix(3, 1), bias: randomFloatMatrix(1, 1)})
	net.AddLayer(ActivationLayer{activation: ReLUMatrix, activationPrime: ReLUMatrix})
	net.SetLoss(Mse, MsePrime)
	net.Fit(x_data, y_labels, 1000, 0.01)

}