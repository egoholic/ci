package stage

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type Stage struct {
	name     string
	commands []*exec.Cmd // TODO: replace with custom command
	stdOut   io.ReadCloser
	stdErr   io.ReadCloser
}

func New(name string, duration time.Duration) *Stage {
	return &Stage{name, nil, nil, nil}
}

func (stage *Stage) AddCommand(ctx context.Context, name string, arg ...string) {
	stage.commands = append(stage.commands, exec.CommandContext(ctx, name, arg...))
}

func (stage *Stage) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for i, cmd := range stage.commands {
		fmt.Println("-(%d) Executing: `%#v`", i, cmd)
		cmd.Run()
	}
}
