package recursion

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRecursion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Recursion Suite")
}
