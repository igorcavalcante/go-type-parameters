package funcs

func SumOfIntegers(xs map[string]int) int {
	var acc int
	for _, v := range xs {
		acc += v
	}

	return acc
}

func SumOfFloats(xs map[string]float64) float64 {
	var acc float64
	for _, v := range xs {
		acc += v
	}

	return acc
}

func SumOfNumbers[K comparable, V float64 | int](xs map[K]V) V {
	var acc V
	for _, v := range xs {
		acc += v
	}

	return acc
}

func SumOfNumbersUsingInterface[K comparable, V Numbers](xs map[K]V) V {
	var acc V
	for _, v := range xs {
		acc += v
	}

	return acc
}

type Numbers interface {
	int | float64 | float32
}
