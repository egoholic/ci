package build_test

import (
	. "github.com/egoholic/ci/build"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CI build", func() {
	Describe("New()", func() {
		It("creates new build", func() {

		})
	})

	Describe(".AddStage()", func() {
		Context("when stage with given name already exists", func() {
			It("fails", func() {

			})
		})

		Context("when stage with given name does not exist", func() {
			It("adds new stage", func() {

			})
		})
	})

	Describe(".Build()", func() {
		Context("when successfull", func() {
			It("prints the output and executes all the commands", func() {

			})
		})

		Context("when failed", func() {
			It("prints the output and breaks commands execution with failure", func() {

			})
		})
	})

	Describe(".Info()", func() {
		Context("when not runned", func() {

		})

		Context("when runned", func() {

		})

		Context("when finished with succeed", func() {

		})

		Context("when finished with failure", func() {

		})
	})
})
