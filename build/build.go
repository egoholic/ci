package build

type Settings struct{}

func (_ *Settings) DBPath() string {
	return "./.DATA/builds"
}

var config = &Settings{}

type Build struct {
	num  int
	ref  string
	path string
}

type Repo struct {
	builds []*Build
}

func New(num int, ref string, path string) *Build {
	return &Build{num, ref, path}
}
