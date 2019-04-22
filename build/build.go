package build

import (
	"fmt"

	"github.com/egoholic/ci/stage"
	"github.com/egoholic/serror"
)

type Build struct {
	buildableName string
	scriptPath    string
	number        int
	stages        map[string]*stage.Stage
	stagesOrder   []string
	currentStage  int
}

func New(buildFile string) *Build {
	return &Build{}
}

func (build *Build) Name() string {
	return fmt.Sprintf("%s-%d", build.buildableName, build.number)
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

func (build *Build) Info() string {
	return ""
}

func (build *Build) Build() (err error) {
	return
}
