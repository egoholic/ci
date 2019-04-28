package repository

type Config struct{}

var config *Config

func (_ *Config) PipelinesDataPath() string {
	return "./data/pipelines"
}

type Repository struct {
	title string
}

func New(title string) *Repository {
	return &Repository{title}
}

func (repo *Repository) Load() (err error) {
	return
}

func (repo *Repository) All() []interface{} {
	return nil
}

func (repo *Repository) Add(pipeline *Pipeline) (err error) {
	return
}
