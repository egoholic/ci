package stage

import "github.com/egoholic/ci/command"

type Stage struct {
	name     string
	commands []*command.Command
}

func New(name string) *Stage {
	return &Stage{name, nil}
}
