package build

type Settings struct{}

func (_ *Settings) DBPath() string {
	return "./.DATA/builds"
}

var config = &Settings{}

type Build struct {
	num int
	ref string
}

type Repo struct {
	builds []*Build
}

func NewRepo(DBPath string) *Repo {

}

func Create(ref string) *Build {

}
