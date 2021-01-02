package CNN

type ActivationLayer struct {
	input Matrix
	activation func(Matrix) Matrix
	activationPrime func(Matrix) Matrix
}

func (a ActivationLayer) ForwardPass (input Matrix) Matrix {
	a.input = input
	return a.activation(input)
}


func (a ActivationLayer) BackPropMatrix (outputError Matrix, alpha float64) Matrix {
	return a.activationPrime(a.input).multiplyMatrix(outputError)
}
