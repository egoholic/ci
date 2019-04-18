package command

type Command struct {
	role    string
	command string
	params  string
}

func New(role, command, params string) *Command {
	return &Command{role, command, params}
}

func (command *Command) Execute() int {
	return 0 // TODO: implement
}
