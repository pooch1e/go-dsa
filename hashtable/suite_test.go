package hashtable

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHashtable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hashtable Suite")
}
