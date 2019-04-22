package build

import (
	"fmt"
	"time"

	"github.com/egoholic/ci/stage"
	"github.com/egoholic/serror"
)

type Build struct {
	scripPath     string
	buildableName string
	uuid          string
	stages        map[string]*stage.Stage
	stagesOrder   []string
	currentStage  int
	startedAt     time.Time
	finishedAt    time.Time
}

func New() *Build {
	return &Build{}
}

func (build *Build) Name() string {
	return fmt.Sprintf("%s-%s", build.buildableName, build.uuid)
}

func (build *Build) AddStage(buildStage *stage.Stage) (err error) {
	stageName := buildStage.Name()
	_, ok := build.stages[stageName]
	if ok {
		err = serror.New(fmt.Sprintf("can't add stage `%s`", stageName), fmt.Sprintf("build `%s` already has stage named `%s`", build.Name(), stageName))
		return
	}
	build.stages[stageName] = buildStage
	build.stagesOrder = append(build.stagesOrder, stageName)
	return
}
