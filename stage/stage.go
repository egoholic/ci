package stage

import (
	"context"

	"github.com/egoholic/ci/cmd"
)

type Stage struct {
	name     string
	commands []*cmd.Command
}

func New(name string) *Stage {
	commands := make([]*cmd.Command, 20)
	return &Stage{name, commands}
}

func (stage *Stage) Name() string {
	return stage.name
}

func (stage *Stage) AddCommand(name string, arg ...string) {
	stage.commands = append(stage.commands, cmd.New(name, arg...))
}

func (stage *Stage) Run(ctx context.Context) {

}
