package stage

import (
	"context"
	"io"
	"os/exec"
	"time"

	"github.com/egoholic/ci/cmd"
)

type Stage struct {
	name       string
	commands   []cmd.Command // TODO: replace with custom command
	stdOut     io.ReadCloser
	stdErr     io.ReadCloser
	startedAt  time.Time
	finishedAt time.Time
}

func New(name string, duration time.Duration) *Stage {
	return &Stage{name, nil, nil, nil}
}

func (stage *Stage) Name() string {
	return stage.name
}

func (stage *Stage) AddCommand(ctx context.Context, name string, arg ...string) {
	stage.commands = append(stage.commands, exec.CommandContext(ctx, name, arg...))
}
