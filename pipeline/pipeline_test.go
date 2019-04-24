package pipeline_test

import (
	. "github.com/egoholic/ci/pipeline"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestWriter struct {
	data []byte
}

func (writer *TestWriter) Write(p []byte) (n int, err error) {
	writer.data = append(writer.data, p...)
	n = len(p)
	return
}

func (writer *TestWriter) String() string {
	return string(writer.data)
}

var _ = Describe("Pipeline", func() {
	Describe("New()", func() {
		It("returns pipeline", func() {
			pipeline := New("Test-Pipeline")
			Expect(pipeline).To(BeAssignableToTypeOf(&Pipeline{}))
			Expect(pipeline.Title()).To(Equal("Test-Pipeline"))
		})
	})

	Describe("NewWithConfig()", func() {
		Context("when correct configuration", func() {
			It("returns pipeline", func() {
				pipeline := NewWithConfig(&PipelineConfig{"Test-Pipeline", nil})
				Expect(pipeline).To(BeAssignableToTypeOf(&Pipeline{}))
				Expect(pipeline.Title()).To(Equal("Test-Pipeline"))
			})
		})

		Context("when wrong configuration", func() {

		})
	})

	Describe("ParseConfig()", func() {
		Context("when correct config source", func() {
			// It("returns pipeline configuration", func() {
			// 	path, _ := filepath.Abs("./test/pipeline.yml")
			// 	source, _ := os.Open(path)
			// 	defer source.Close()

			// 	config, err := ParseConfig(source)
			// 	Expect(err).ShouldNot(HaveOccurred())
			// 	Expect(config.Title).To(Equal("Pipeline1"))
			// })
		})
	})

	Describe("NewStage()", func() {
		It("returns stage", func() {
			Expect(NewStage("Stage1")).To(BeAssignableToTypeOf(&Stage{}))
		})
	})

	Describe("Stage", func() {
		Describe(".Title()", func() {
			It("returns title", func() {
				stage := NewStage("Stage1")
				Expect(stage.Title()).To(Equal("Stage1"))
			})
		})

		Describe(".AddCommand()", func() {
			It("adds new command", func() {
				stage := NewStage("Stage1")
				err := stage.AddCommand("go", "version")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Describe(".Run()", func() {
			Context("when everything performs successfully", func() {

			})

			Context("when some command fails", func() {

			})
		})
	})

	Describe("Pipeline", func() {
		Describe(".Title()", func() {
			It("returns title", func() {
				pipeline := New("Pipeline1")
				Expect(pipeline.Title()).To(Equal("Pipeline1"))
			})
		})

		Describe(".Run()", func() {
			Context("when everything performs successfully", func() {
				It("succeed", func() {

				})
			})

			Context("when some command fails", func() {
				It("fails", func() {

				})
			})
		})
	})
})
