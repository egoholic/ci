package pipeline_test

import (
	. "github.com/egoholic/ci/pipeline"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pipeline", func() {
	Describe("New()", func() {
		Context("when correct configuration source", func() {
			It("returns pipeline", func() {
				pipeline, err := New("./test/pipeline.yml")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(pipeline).To(BeAssignableToTypeOf(&Pipeline{}))
				Expect(pipeline.Title()).To(Equal("Test-Pipeline"))
			})
		})
	})

	Describe(".Run()", func() {

	})
})
