package set

func Intersect[T comparable](lists ...[]T) (out []T) {
	hist := make(map[T]int)
	for i, list := range lists {
		for _, item := range list {
			if hist[item] == i {
				hist[item]++
			}
		}
	}

	for item, count := range hist {
		if count == len(lists) {
			out = append(out, item)
		}
	}

	return
}

func Union[T comparable](lists ...[]T) (out []T) {
	hist := make(map[T]struct{})
	for _, list := range lists {
		for _, item := range list {
			hist[item] = struct{}{}
		}
	}

	for item := range hist {
		out = append(out, item)
	}
	return
}

func Excluding[T comparable](list []T, listOfValuesToExclude ...[]T) (out []T) {
	hist := make(map[T]bool)
	for _, item := range list {
		hist[item] = true
	}
	for _, valuesToExclude := range listOfValuesToExclude {
		for _, valueToExclude := range valuesToExclude {
			if hist[valueToExclude] {
				delete(hist, valueToExclude)
			}
		}
	}

	for item := range hist {
		out = append(out, item)
	}

	return
}
