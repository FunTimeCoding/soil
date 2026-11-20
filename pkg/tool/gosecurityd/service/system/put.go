package system

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (s *Service) Put(
	v *system.System,
	packages []installed.Installed,
) error {
	if !constant.SystemNamePattern.MatchString(v.Name) {
		return validation.New("system name malformed: %s", v.Name)
	}

	if v.Release == "" {
		return validation.New("system release missing: %s", v.Name)
	}

	if len(packages) == 0 && s.debian.Covering(v.Release) != nil {
		return validation.New(
			"system without packages on a tracked release: %s",
			v.Name,
		)
	}

	for _, p := range packages {
		if p.Name == "" || p.Version == "" {
			return validation.New(
				"package without name or version on system: %s",
				v.Name,
			)
		}

		if p.Source == "" || p.SourceVersion == "" {
			return validation.New(
				"package %s without source name or version on system: %s",
				p.Name,
				v.Name,
			)
		}
	}

	return s.store.PutSystem(v, packages)
}
