package agent

import (
	"github.com/egoholic/ci/cmd"
)

type Agent struct {
	title string
}

func (agent *Agent) Title() string {
	return agent.title
}

func (agent *Agent) Run(command *cmd.Command) string {

}
