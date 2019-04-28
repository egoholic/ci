package stage

import (
	"context"

	"github.com/egoholic/ci/cmd"
)

type Config struct {
	Title    string
	Commands []string `yaml: ",flow"`
}

type Stage struct {
	title    string
	commands []*cmd.Command
}

func New(name string) *Stage {
	return &Stage{name, nil}
}

func (stage *Stage) Title() string {
	return stage.title
}

func (stage *Stage) AddCommand(name string, arg ...string) error {
	stage.commands = append(stage.commands, cmd.New(name, arg...))
	return nil
}

func (stage *Stage) Run(ctx context.Context) {

}
