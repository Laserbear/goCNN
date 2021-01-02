package CNN

type FullyConnectedLayer struct {
	weights Matrix
	input Matrix
	bias Matrix
}


func (fc FullyConnectedLayer) ForwardPass (input Matrix) Matrix {
	fc.input = input
	return fc.weights.multiplyMatrix(input).add(fc.bias)
}


func (fc FullyConnectedLayer) BackPropMatrix (outputError Matrix, alpha float64) Matrix {
	inputError := outputError.multiplyMatrix(fc.weights.transpose())
	weightsError := fc.input.transpose().multiplyMatrix(outputError)

	fc.weights = fc.weights.add(weightsError.multiplyFloat(-1 * alpha))
	fc.bias = fc.bias.add(outputError.multiplyFloat(-1 * alpha))

	return inputError
}



