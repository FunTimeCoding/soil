package host_command

type HostCommand struct {
	Group     string `yaml:"group"`
	Module    string `yaml:"module"`
	Arguments string `yaml:"arguments"`
}
