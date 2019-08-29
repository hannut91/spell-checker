package differ

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiffer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Differ Suite")
}
