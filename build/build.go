package build

import "time"

type Build struct {
	startedAt       time.Time
	builtAt         time.Time
	reference       string
	currentStage    string
	messagesTracked map[string][]string
}

func New() *Build {
	return &Build{}
}

func (build *Build) JSON() string {
	return "" // TODO: implement
}
