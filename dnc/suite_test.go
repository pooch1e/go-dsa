package dnc

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDnc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dnc Suite")
}
