package pipeline

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/egoholic/ci/stage"
	yaml "gopkg.in/yaml.v3"
)

type StageConfig struct {
	Title    string   `yaml:"title"`
	Commands []string `yaml:"commands"`
}

type PipelineConfig struct {
	Title  string        `yaml:"title"`
	Stages []StageConfig `yaml:"stages"`
}

type Pipeline struct {
	title        string
	configSource string
	stages       map[string]*stage.Stage
	stagesOrder  []string
	currentStage int
}

func New(pipelineFile string) (pipeline *Pipeline, err error) {
	var (
		file           *os.File
		data           []byte
		pipelineConfig *PipelineConfig
	)

	file, err = os.Open(pipelineFile)
	if err != nil {
		fmt.Printf("ERROR in builder: %s", err.Error())
		return
	}
	defer file.Close()

	for {
		_, err = file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("ERROR in builder: %s", err.Error())
			return
		}
	}

	err = yaml.Unmarshal(data, pipelineConfig)
	if err != nil {
		return
	}
	pipeline = &Pipeline{pipelineConfig.Title, pipelineFile, nil, nil, 0}
	return
}

func (pipeline *Pipeline) Title() string {
	return pipeline.title
}

func (pipeline *Pipeline) Run(ctx context.Context) {

}
