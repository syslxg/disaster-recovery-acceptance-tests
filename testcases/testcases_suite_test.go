package testcases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestcases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testcases Suite")
}
