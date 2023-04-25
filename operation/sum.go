package operation

// Sum //Suma de x cantidad de enteros
func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
