package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPetitBacApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PetitBacApp Suite")
}
