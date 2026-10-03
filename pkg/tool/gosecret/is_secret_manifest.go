package gosecret

import (
	"go.yaml.in/yaml/v3"
	"os"
)

func IsSecretManifest(path string) (bool, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return false, e
	}

	var m SecretManifest

	if e := yaml.Unmarshal(b, &m); e != nil {
		return false, nil
	}

	return m.Kind == "Secret", nil
}
