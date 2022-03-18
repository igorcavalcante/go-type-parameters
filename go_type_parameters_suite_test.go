package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoTypeParameters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTypeParameters Suite")
}
