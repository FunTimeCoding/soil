package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

func (s *Service) checkStrictServer(
	v *virtual_file_system.System,
	path string,
) {
	specFile := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(specFile) {
		return
	}

	configurationFile := filepath.Join(
		path,
		"generated",
		"server",
		"config.yaml",
	)

	if !v.Has(configurationFile) {
		s.addConcern(
			constant.MissingStrictServerKey,
			constant.MissingStrictServerText,
			path,
		)

		return
	}

	var c strictServerConfiguration

	if yaml.Unmarshal(v.Read(configurationFile), &c) != nil {
		s.addConcern(
			constant.MissingStrictServerKey,
			constant.MissingStrictServerText,
			path,
		)

		return
	}

	if !c.Generate.StrictServer {
		s.addConcern(
			constant.MissingStrictServerKey,
			constant.MissingStrictServerText,
			path,
		)
	}
}
