package math 

func Average(arr []float64) float64 {
	total := float64(0)
	for _, element := range arr {
		total += element
	}
	return total/float64(len(arr))
}