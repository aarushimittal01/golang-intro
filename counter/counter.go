package counter

func Get_counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Get_value() (func() int, func() int) {
	value := 0

	up := func() int {
		value++
		return value
	}

	down := func() int {
		value--
		return value
	}

	return up, down
}

func Get_series(offset, difference int) func() int {
	return func() int {
		offset += difference
		return offset
	}
}

//passing condition in the function
func Get_numbers(predicate func(int) bool) func() int {
	counter := 0
	return func() int {
		for {
			counter++
			if predicate(counter) {
				return counter
			}
		}
	}
}
