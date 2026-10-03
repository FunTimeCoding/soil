package installed

import (
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func versions(directory string) map[string]string {
	result := make(map[string]string)
	r := run.New()
	r.Panic = false
	r.Directory = directory
	output := r.Start(
		libraryConstant.Go,
		libraryConstant.List,
		libraryConstant.ModuleArgument,
		constant.FormatArgument,
		constant.ModuleListTemplate,
		constant.AllModules,
	)

	if r.Error != nil {
		return result
	}

	for _, line := range strings.Split(output, stringsConstant.Unix) {
		if fields := strings.Fields(line); len(fields) == 2 {
			result[fields[0]] = fields[1]
		}
	}

	return result
}
