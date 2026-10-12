package gosecret

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"go.yaml.in/yaml/v3"
	"os"
)

func readManifest(path string) (*SecretManifest, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return nil, fmt.Errorf("read manifest: %w", e)
	}

	var result SecretManifest

	if f := yaml.Unmarshal(b, &result); f != nil {
		return nil, fmt.Errorf("parse manifest: %w", f)
	}

	if result.Kind != "Secret" {
		return nil, validation.New("not a Secret manifest: %s", path)
	}

	if result.Metadata == nil || result.Metadata.Name == "" {
		return nil, validation.New("Secret manifest has no name: %s", path)
	}

	return &result, nil
}
