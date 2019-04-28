package pipeline

import (
	"context"
	"io"

	"github.com/egoholic/ci/pipeline/stage"
	yaml "gopkg.in/yaml.v2"
)

type PipelineConfig struct {
	Title  string
	Stages []stage.StageConfig `yaml: ",flow"`
}

type Pipeline struct {
	title     string
	stages    []*stage.Stage
	stagesCur int
}

func ParseYAMLConfig(source io.Reader) (config *PipelineConfig, err error) {
	data := make([]byte, 4098)
	n, err := source.Read(data)
	if err != nil {
		return
	}
	data = data[0 : n-1]
	err = yaml.Unmarshal(data, config)
	return
}

func NewWithConfig(config *PipelineConfig) *Pipeline {
	return &Pipeline{config.Title, nil, 0}
}

func New(title string) *Pipeline {
	return &Pipeline{title, nil, 0}
}

func (pipeline *Pipeline) Title() string {
	return pipeline.title
}

func (pipeline *Pipeline) Run(ctx context.Context) (err error) {
	return
}

func (pipeline *Pipeline) ReRun(ctx context.Context) (err error) {
	return
}

func (pipeline *Pipeline) Stop() (err error) {
	return
}
