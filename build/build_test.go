package build_test

import (
	. "github.com/egoholic/ci/build"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CI build", func() {
	Describe("New()", func() {
		It("returns build", func() {
			Expect(New(1, "abcdef12345", "/builds")).To(BeAssignableToTypeOf(&Build{}))
		})
	})
})
