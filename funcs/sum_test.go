package funcs_test

import (
	"github.com/igorcavalcante/go-type-parameters/funcs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum examples", func() {

	ints := map[string]int{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	var _ = Describe("Sum of Integers", func() {

		When("Calculating", func() {
			It("should return the sum of all values", func() {
				Expect(funcs.SumOfIntegers(ints)).To(Equal(46))
			})
		})

	})

	var _ = Describe("Sum of floaters", func() {

		When("calculating", func() {
			It("should return the sum of all values", func() {
				Expect(funcs.SumOfFloats(floats)).To(Equal(62.97))
			})
		})

	})

	var _ = Describe("Generic Sum", func() {

		When("calculating integers", func() {
			It("should return the sum of all values", func() {
				Expect(funcs.SumOfNumbers(ints)).To(Equal(46))
				Expect(funcs.SumOfNumbersUsingInterface(ints)).To(Equal(46))
			})
		})

		When("calculating floats", func() {
			It("should return the sum of all values", func() {
				Expect(funcs.SumOfNumbers(floats)).To(Equal(62.97))
				Expect(funcs.SumOfNumbersUsingInterface(floats)).To(Equal(62.97))
			})
		})

	})

})
