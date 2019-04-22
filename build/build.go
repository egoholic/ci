package build

import (
	"fmt"
	"io"
	"os"

	"github.com/egoholic/ci/stage"
	"github.com/egoholic/serror"
	yaml "gopkg.in/yaml.v3"
)

type StageConfig struct {
	Title    string   `yaml:"title"`
	Commands []string `yaml:"commands"`
}
type BuildConfig struct {
	Title  string        `yaml:"title"`
	Stages []StageConfig `yaml:"stages"`
}

type Build struct {
	title        string
	num          int
	scriptPath   string
	stages       map[string]*stage.Stage
	stagesOrder  []string
	currentStage int
}

func New(buildFile string) (build *Build, err error) {
	var (
		file        *os.File
		data        []byte
		buildConfig *BuildConfig
	)

	file, err = os.Open(buildFile)
	if err != nil {
		fmt.Printf("ERROR in builder: %s", err.Error())
		return
	}

	for {
		_, err = file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("ERROR in builder: %s", err.Error())
		}
	}

	err = yaml.Unmarshal(data, buildConfig)
	if err != nil {
		return
	}
	build = &Build{buildConfig.Title, 1, buildFile, nil, nil, 0}
	return
}

func (build *Build) Name() string {
	return fmt.Sprintf("%s-%d", build.title, build.num)
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
