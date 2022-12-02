package math

func Min(ns ...int) int {
	min := ns[0]
	for _, n := range ns {
		if n < min {
			min = n
		}
	}
	return min
}
func Max(ns ...int) int {
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}
func Extremum(n ...int) (int, int) {
	return Min(n...), Max(n...)
}
func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
