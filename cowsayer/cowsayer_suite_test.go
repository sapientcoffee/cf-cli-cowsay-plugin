package cowsayer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCowsayer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cowsayer Suite")
}
