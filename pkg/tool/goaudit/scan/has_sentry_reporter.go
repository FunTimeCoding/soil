package scan

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func hasSentryReporter(
	v *virtual_file_system.System,
	directory string,
) bool {
	for _, name := range v.MustReadDirectory(directory) {
		if !strings.HasSuffix(name, constant.GoExtension) {
			continue
		}

		f, _, e := parse.Source(
			name,
			v.ReadString(filepath.Join(directory, name)),
		)

		if e != nil {
			continue
		}

		for _, name := range []string{"New", "NewOptional"} {
			if parse.HasCall(
				f,
				"github.com/funtimecoding/soil/pkg/errors/sentry/reporter",
				name,
			) {
				return true
			}
		}
	}

	return false
}
