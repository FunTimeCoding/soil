package instance

type Instance struct {
	Name     string  `yaml:"name"`
	Host     string  `yaml:"host"`
	Token    string  `yaml:"token"`
	Groups   []int64 `yaml:"groups"`
	Projects []int64 `yaml:"projects"`
}
