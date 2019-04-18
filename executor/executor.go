package executor

import "github.com/egoholic/ci/stage"

type Executor struct {
	stages []*stage.Stage
}

func New() *Executor {
	return &Executor{}
}

func (executor *Executor) AddStage(buildStage *stage.Stage) {
	executor.stages = append(executor.stages, buildStage)
}
