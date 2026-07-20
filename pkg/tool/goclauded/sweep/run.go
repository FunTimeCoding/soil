package sweep

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"path/filepath"
)

func Run(harbor string) *Result {
	system.MakeDirectory(harbor)
	sources := collectSources()
	r := &Result{}

	for _, source := range sources {
		name := filepath.Base(source)
		destination := filepath.Join(harbor, name)

		switch synchronize(source, destination) {
		case constant.ActionSkip:
			r.Skipped++
		case constant.ActionCopy:
			copyFile(source, destination)
			r.Copied++
		case constant.ActionUpdate:
			copyFile(source, destination)
			r.Updated++
		}
	}

	return r
}
