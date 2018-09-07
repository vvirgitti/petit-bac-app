package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bac", func() {

	Describe("Random letter", func () {
		It("should display a random letter", func() {
			Expect(generateLetter())
		})
	})
})
