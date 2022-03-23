package types_test

import (
	"github.com/igorcavalcante/go-type-parameters/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree", func() {

	var intTree types.Tree[int]
	var stringTree types.Tree[string]

	BeforeEach(func() {
		intTree = types.Tree[int]{Value: 42}
		stringTree = types.Tree[string]{Value: "foo"}
	})

	Describe("accepting different types", func() {
		When("the tree is composed by strings", func() {
			It("should expose the string as the value", func() {
				result := stringTree.GetValue()
				Expect(result).To(BeIdenticalTo("foo"))
			})
		})

		When("the tree is composed by integers", func() {
			It("should expose the number as the value", func() {
				result := intTree.GetValue()
				Expect(result).To(BeIdenticalTo(42))
			})
		})
	})

	Describe("counting", func() {
		When("the tree has just one element", func() {
			It("should return 1", func() {
				Expect(intTree.Count()).To(Equal(1))
			})
		})

		When("the tree has three elements", func() {
			BeforeEach(func() {
				intTree.Left = &types.Tree[int]{Value: 2}
				intTree.Right = &types.Tree[int]{Value: 1}
			})

			It("Should should return 3 as counted elements", func() {
				Expect(intTree.Count()).To(Equal(3))
			})
		})
	})
})
