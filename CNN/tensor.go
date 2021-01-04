package CNN

type tensor struct {
	rawValues  []float64
	dimensions []int
	length     uint
}

func (t1 tensor) hadamard_product(t2 tensor) tensor { //only works if they are same dimensions rn
	productTensor := make([]float64, t1.length)
	for i := uint(0); i < t1.length; i++ {
		productTensor[i] = t1.rawValues[i] * t2.rawValues[i]
	}
	return tensor{rawValues: []float64(productTensor), dimensions: t1.dimensions, length: t1.length}
}
