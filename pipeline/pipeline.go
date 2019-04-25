package pipeline

import (
	"context"
	"io"

	"github.com/egoholic/ci/cmd"
	yaml "gopkg.in/yaml.v2"
)

type StageConfig struct {
	Title    string
	Commands []string `yaml: ",flow"`
}

type Stage struct {
	title    string
	commands []*cmd.Command
}

func NewStage(name string) *Stage {
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

type PipelineConfig struct {
	Title  string
	Stages []StageConfig `yaml: ",flow"`
}

type Pipeline struct {
	title     string
	stages    []*Stage
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

type RepoConfig struct{}

var config *RepoConfig

func (_ *RepoConfig) PipelinesDataPath() string {
	return "./data/pipelines"
}

type Repo struct {
	title string
}

func NewRepo(title string) *Repo {
	return &Repo{title}
}

func (repo *Repo) Load() (err error) {
	return
}

func (repo *Repo) All() []interface{} {
	return nil
}

func (repo *Repo) Add(pipeline *Pipeline) (err error) {

}
