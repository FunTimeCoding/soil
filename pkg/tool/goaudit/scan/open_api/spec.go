package open_api

type Spec struct {
	Header Header         `yaml:"info"`
	Paths  map[string]any `yaml:"paths"`
}
