package cmd

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	name string
	args []string
}

func New(name string, args ...string) *Command {
	return &Command{name, args}
}

func (command *Command) String() string {
	builder := &strings.Builder{}
	builder.WriteString(command.name)
	for _, arg := range command.args {
		builder.WriteRune(' ')
		builder.WriteString(arg)
	}
	return builder.String()
}

func (command *Command) Run(ctx context.Context, errWriter io.Writer, outWriter io.Writer) (err error) {
	cmd := exec.CommandContext(ctx, command.name, command.args...)
	cmd.Stderr = errWriter
	cmd.Stdout = outWriter
	err = cmd.Run()
	return
}

type Writer struct {
	path string
	file *os.File
}

func NewWriter(path string) (writer *Writer, err error) {
	var file *os.File
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err = os.Create(path)
		if err != nil {
			return
		}
	}
	file, err = os.Open(path)
	if err != nil {
		return
	}
	writer = &Writer{path, file}
	return
}

func (writer *Writer) Close() (err error) {
	return writer.file.Close()
}

func (writer *Writer) Write(p []byte) (n int, err error) {
	return writer.file.Write(p)
}
