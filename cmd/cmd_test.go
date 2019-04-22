package cmd_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	. "github.com/egoholic/ci/cmd"
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

var _ = Describe("cmd", func() {
	Describe("Command", func() {
		Describe("New()", func() {
			It("returns command", func() {
				Expect(New("go", "version")).To(BeAssignableToTypeOf(&Command{}))
			})
		})

		Describe(".String()", func() {
			It("returns command's string representation", func() {
				command := New("go", "version")
				Expect(command.String()).To(Equal("go version"))
			})
		})

		Describe(".Run()", func() {
			Context("when correct command", func() {
				It("succeed", func() {
					command := New("go", "version")
					outWriter := &TestWriter{}
					errWriter := &TestWriter{}
					ctx := context.TODO()
					err := command.Run(ctx, errWriter, outWriter)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(errWriter.String()).To(BeEmpty())
					Expect(outWriter.String()).To(Equal("go version go1.11 darwin/amd64\n"))
				})
			})

			Context("when wrong command", func() {
				It("fails", func() {
					command := New("ololo", "version")
					outWriter := &TestWriter{}
					errWriter := &TestWriter{}
					ctx := context.TODO()
					err := command.Run(ctx, errWriter, outWriter)
					Expect(err).Should(HaveOccurred())
					Expect(errWriter.String()).To(Equal(""))
					Expect(outWriter.String()).To(BeEmpty())
				})
			})
		})
	})

	Describe("Writer", func() {
		Describe("NewWriter()", func() {
			Context("when correct path", func() {
				It("creates new writer", func() {
					Expect(NewWriter("./testDir/out")).To(BeAssignableToTypeOf(&Writer{}))
				})
			})
		})

		Describe(".Write()", func() {
			Context("when writer is open", func() {
				It("writes", func() {
					path, err := filepath.Abs("./testDir/out")
					Expect(err).NotTo(HaveOccurred())
					fmt.Printf(path)
					defer func() {
						Expect(err).NotTo(HaveOccurred())
						file, err := os.Create(path)
						Expect(err).NotTo(HaveOccurred())
						defer file.Close()
					}()
					defer os.Remove(path)

					writer, err := NewWriter(path)
					Expect(err).ShouldNot(HaveOccurred())
					_, err = writer.Write([]byte("test"))
					Expect(err).ShouldNot(HaveOccurred())
					err = writer.Close()
					Expect(err).ShouldNot(HaveOccurred())
					file, err := os.Open(path)
					Expect(err).ShouldNot(HaveOccurred())
					content := make([]byte, len([]byte("test")))
					_, err = file.Read(content)
					if err != nil && err != io.EOF {
						Expect(err).ShouldNot(HaveOccurred())
					}

					//Expect(len(content)).To(Equal(n))
					Expect(content).To(Equal([]byte("test")))
				})
			})

			Context("when writer is closed", func() {
				It("fails", func() {
					path, err := filepath.Abs("./testDir/out")
					Expect(err).NotTo(HaveOccurred())

					fmt.Printf(path)
					defer func() {
						file, err := os.Create(path)
						Expect(err).NotTo(HaveOccurred())
						defer file.Close()
					}()
					defer os.Remove(path)

					writer, err := NewWriter(path)
					Expect(err).ShouldNot(HaveOccurred())
					err = writer.Close()
					Expect(err).ShouldNot(HaveOccurred())
					n, err := writer.Write([]byte("test"))
					Expect(err).To(HaveOccurred())
					Expect(n).To(BeZero())
				})
			})
		})

		Describe(".Close()", func() {
			It("closes", func() {
				path, err := filepath.Abs("./testDir/out")
				Expect(err).NotTo(HaveOccurred())
				fmt.Printf(path)
				defer func() {
					file, err := os.Create(path)
					Expect(err).NotTo(HaveOccurred())
					defer file.Close()
				}()
				defer os.Remove(path)

				writer, err := NewWriter(path)
				Expect(err).ShouldNot(HaveOccurred())
				writer.Write([]byte("test"))
				err = writer.Close()
				Expect(err).ShouldNot(HaveOccurred())
				err = writer.Close()
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
