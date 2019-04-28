package record

import (
	"github.com/egoholic/ci/pipeline/resolution"
)

type Record struct {
	title      string
	logs       map[string]string
	resolution *resolution.Resolution
}

func New(title string) *Record {
	return &Record{title, nil, WAITING}
}
