package funcs_test

import (
	"github.com/igorcavalcante/go-type-parameters/funcs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("My example functions", func() {

	integers := []int{1, 2, 3}

	var _ = Describe("My map function", func() {

		When("mapping a square function over a slice of integers", func() {

			square := func(x int) int { return x * x }
			result := []int{1, 4, 9}

			It("should return a slice containing the square of all original elements", func() {
				Expect(funcs.GMap(integers, square)).To(Equal(result))
			})
		})

	})

	var _ = Describe("My Reduce function", func() {

		When("reducing a slice of integers with a sum function", func() {
			sum := func(x int, y int) int { return x + y }
			result := 6

			It("should return the sum of all elements", func() {
				Expect(funcs.GReduce(integers, 0, sum)).To(Equal(result))
			})
		})

	})

	var _ = Describe("My Find function", func() {

		When("Finding the elements greater than 2", func() {
			predicate := func(x int) bool { return x > 2 }
			result := []int{3}

			It("should return a slice with one element: [3]", func() {
				Expect(funcs.GFind(integers, predicate)).To(Equal(result))
			})
		})

		When("Finding the elements greater than 3", func() {
			predicate := func(x int) bool { return x > 3 }
			var result []int

			It("should return an empty slice", func() {
				Expect(funcs.GFind(integers, predicate)).To(Equal(result))
			})
		})

	})

})
