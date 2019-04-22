package stage_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stage Suite")
}
