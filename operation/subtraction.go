package operation

// Sub //resta de x cantidad de enteros
func Sub(values ...int) int {
	total := 0
	for _, value := range values {
		total -= value
	}
	return total
}
