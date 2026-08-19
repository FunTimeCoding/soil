package gosecret

import (
	"gopkg.in/yaml.v3"
	"os"
)

func IsSecretManifest(path string) (bool, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return false, e
	}

	var m SecretManifest

	if e := yaml.Unmarshal(b, &m); e != nil {
		// Not a valid YAML or doesn't match our structure, skip
		return false, nil
	}

	return m.Kind == "Secret", nil
}
