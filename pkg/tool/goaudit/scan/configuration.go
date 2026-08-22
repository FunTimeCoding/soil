package scan

type Configuration struct {
	Exclude      []string                       `yaml:"exclude"`
	Suppress     map[string]map[string][]string `yaml:"suppress"`
	ModelContext map[string]string              `yaml:"model_context"`
	Aliases      map[string]string              `yaml:"model_context_alias"`
}
