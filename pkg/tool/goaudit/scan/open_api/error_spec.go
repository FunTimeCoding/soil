package open_api

type ErrorSpec struct {
	Paths      map[string]map[string]Operation `yaml:"paths"`
	Components Components                      `yaml:"components"`
}
