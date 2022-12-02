package set

func Intersect(lists ...[]rune) (out []rune) {
	hist := make(map[rune]int)
	for _, list := range lists {
		for _, letter := range list {
			hist[letter]++
		}
	}
	for letter, count := range hist {
		if count == len(lists) {
			out = append(out, letter)
		}
	}
	return
}

func Union(lists ...[]rune) (out []rune) {
	hist := make(map[rune]struct{})
	for _, list := range lists {
		for _, letter := range list {
			hist[letter] = struct{}{}
		}
	}

	for letter := range hist {
		out = append(out, letter)
	}
	return
}

func Difference(lists ...[]rune) (out []rune) {
	hist := make(map[rune]int)
	for _, list := range lists {
		for _, letter := range list {
			hist[letter]++
		}
	}
	for letter, count := range hist {
		if count == 1 {
			out = append(out, letter)
		}
	}
	return
}

func Excluding(list []rune, listOfValuesToExclude ...[]rune) (out []rune) {
	hist := make(map[rune]bool)
	for _, letter := range list {
		hist[letter] = true
	}
	for _, valuesToExclude := range listOfValuesToExclude {
		for _, valueToExclude := range valuesToExclude {
			if hist[valueToExclude] {
				delete(hist, valueToExclude)
			}
		}
	}

	for letter := range hist {
		out = append(out, letter)
	}

	return
}
