package math

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](ns ...T) T {
	min := ns[0]
	for _, n := range ns {
		if n <= min {
			min = n
		}
	}
	return min
}
func Max[T constraints.Ordered](ns ...T) T {
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}
func Extremum[T constraints.Ordered](n ...T) (T, T) {
	return Min(n...), Max(n...)
}
func Abs[T constraints.Integer](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

// Return:
// -1 if n in negative,
// 0 if n = 0,
// 1 if n is positive
func Sign[T constraints.Integer](n T) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	default:
		return 0
	}
}
