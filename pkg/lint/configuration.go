package lint

type Configuration struct {
	Registries []string            `yaml:"registries"`
	Reflow     ReflowConfiguration `yaml:"reflow"`
}
