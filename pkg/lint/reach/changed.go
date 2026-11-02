package reach

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"strings"
)

func Changed(root string) (map[string]bool, error) {
	r := run.New().NoPanic()
	r.Directory = root
	r.Start(
		gitConstant.Command,
		gitConstant.Status,
		gitConstant.Porcelain,
		constant.NullSeparated,
		constant.AllUntracked,
	)

	if r.Error != nil {
		return nil, r.Error
	}

	result := map[string]bool{}
	entries := strings.Split(r.OutputString, "\x00")

	for i := 0; i < len(entries); i++ {
		entry := entries[i]

		if len(entry) <= constant.PorcelainStatusWidth {
			continue
		}

		result[filepath.Join(root, entry[constant.PorcelainStatusWidth:])] = true

		if entry[0] == 'R' || entry[0] == 'C' {
			i++

			if i < len(entries) && entries[i] != "" {
				result[filepath.Join(root, entries[i])] = true
			}
		}
	}

	return result, nil
}
