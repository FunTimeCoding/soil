package job

type Job struct {
	Paths []string `yaml:"paths"`
	Run   string   `yaml:"run"`
}
