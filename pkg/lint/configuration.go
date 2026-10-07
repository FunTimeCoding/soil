package lint

type configuration struct {
	Registries []string            `yaml:"registries"`
	Reflow     reflowConfiguration `yaml:"reflow"`
}
