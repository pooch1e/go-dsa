package dp

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dp Suite")
}
