package scan

import (
	"go.yaml.in/yaml/v3"
	"os"
)

func LoadConfiguration(path string) *Configuration {
	b, e := os.ReadFile(path)

	if e != nil {
		return &Configuration{}
	}

	var c Configuration

	if e := yaml.Unmarshal(b, &c); e != nil {
		return &Configuration{}
	}

	return &c
}
