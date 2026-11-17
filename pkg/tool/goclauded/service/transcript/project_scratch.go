package transcript

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist/setting"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/scratch"
	"io/fs"
)

func projectScratch(
	cache map[string][]string,
	project string,
) ([]string, error) {
	if roots, okay := cache[project]; okay {
		return roots, nil
	}

	var directories []string

	for _, path := range settingsPaths(project) {
		f, e := setting.Read(path)

		if errors.Is(e, fs.ErrNotExist) {
			continue
		}

		if e != nil {
			return nil, e
		}

		if f.Permissions != nil {
			directories = append(
				directories,
				f.Permissions.AdditionalDirectories...,
			)
		}
	}

	roots := scratch.Roots(project, directories)
	cache[project] = roots

	return roots, nil
}
