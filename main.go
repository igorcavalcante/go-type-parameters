package main

import "fmt"

func main() {
	square := func(x int) int {
		return x * x
	}
	xs := []int{1, 2, 3}
	res := gMap(xs, square)

	fmt.Printf("%v\n", res)

	sum := func(x int, y int) int {
		return x + y
	}

	fmt.Printf("%v\n", gReduce(xs, 0, sum))

	gT2 := func(x int) bool {
		return x > 2
	}

	fmt.Printf("%v\n", gFind(xs, gT2))
}

func gMap[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))
	for i, v := range xs {
		ys[i] = f(v)
	}

	return ys
}

func gReduce[T, U any](xs []T, acc U, f func(U, T) U) U {
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
	}

	return acc
}

func gFind[T comparable](xs []T, f func(T) bool) []T {
	var ys []T
	for _, v := range xs {
		if f(v) {
			ys = append(ys, v)
		}
	}
	return ys
}
