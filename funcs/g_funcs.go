package funcs

func GMap[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))
	for i, v := range xs {
		ys[i] = f(v)
	}

	return ys
}

func GReduce[T, U any](xs []T, acc U, f func(U, T) U) U {
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
	}

	return acc
}

func GFind[T comparable](xs []T, f func(T) bool) []T {
	var ys []T
	for _, v := range xs {
		if f(v) {
			ys = append(ys, v)
		}
	}
	return ys
}
