package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

func openAPIConcern(
	v *virtual_file_system.System,
	path string,
	name string,
) *concern.Concern {
	file := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(file) {
		return nil
	}

	var s openAPISpec

	if yaml.Unmarshal(v.Read(file), &s) != nil {
		return nil
	}

	if s.Header.Title == "" || s.Header.Title == name {
		return nil
	}

	return concern.NewPackage(
		OpenAPIMismatchKey,
		fmt.Sprintf(
			"OpenAPI title %s does not match identity %s",
			s.Header.Title,
			name,
		),
		path,
	)
}
