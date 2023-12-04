package tools

func Sum(v ...int) int {
	total := 0
	for i := 0; i < len(v); i++ {
		total += v[i]
	}
	return total
}
