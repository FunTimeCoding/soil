package instance_tester

type Instance struct {
	Name   string `yaml:"name"`
	Host   string `yaml:"host"`
	Broken bool   `yaml:"broken"`
}
