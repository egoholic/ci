package ci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ci Suite")
}
