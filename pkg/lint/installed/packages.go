package installed

import (
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"strings"
)

func packages(directory string) (map[string]string, map[string][]string) {
	directories := make(map[string]string)
	dependencies := make(map[string][]string)
	r := run.New()
	r.Panic = false
	r.Directory = directory
	output := r.Start(
		libraryConstant.Go,
		libraryConstant.List,
		constant.FormatArgument,
		constant.PackageListTemplate,
		goModConstant.AllPackages,
	)

	if r.Error != nil {
		return directories, dependencies
	}

	for _, line := range strings.Split(output, stringsConstant.Unix) {
		fields := strings.Split(line, stringsConstant.Tab)

		if len(fields) < 3 {
			continue
		}

		relative, e := filepath.Rel(directory, fields[1])

		if e != nil {
			continue
		}

		directories[fields[0]] = filepath.ToSlash(relative)
		dependencies[fields[0]] = strings.Fields(fields[2])
	}

	return directories, dependencies
}
