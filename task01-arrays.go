package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var x float32 = 0
	for _, num := range input {
		x += num
	}
	return x / float32(len(input))
}
