package backtracking

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBacktracking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backtracking Suite")
}
