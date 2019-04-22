package build

import (
	"time"

	"github.com/egoholic/ci/stage"
)

type Build struct {
	uuid         string
	reference    string
	currentStage int
	stages       []*stage.Stage
	startedAt    time.Time
	builtAt      time.Time
	image        string
}

func New() *Build {
	return &Build{}
}

func (build *Build) AddStage(buildStage *stage.Stage) {
	build.stages = append(build.stages, buildStage)
}

func (build *Build) MessagesChannel() chan []string {
	return build.messagesChannel
}
