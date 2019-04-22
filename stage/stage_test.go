package stage_test

import (
	. "github.com/egoholic/ci/stage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stage", func() {
	Describe("New()", func() {
		It("returns stage", func() {
			Expect(New("test-stage")).To(BeAssignableToTypeOf(&Stage{}))
		})
	})

	Describe(".Name()", func() {
		It("returns stage name", func() {
			stage := New("test-stage")
			Expect(stage.Name()).To(Equal("test-stage"))
		})
	})

	Describe(".AddCommand()", func() {
		It("adds command to stage", func() {
			stage := New("test-stage")
			stage.AddCommand("go", "version")
		})
	})
})
