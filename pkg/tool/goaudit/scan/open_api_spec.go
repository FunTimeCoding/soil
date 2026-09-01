package scan

type openAPISpec struct {
	Header openAPIHeader  `yaml:"info"`
	Paths  map[string]any `yaml:"paths"`
}
