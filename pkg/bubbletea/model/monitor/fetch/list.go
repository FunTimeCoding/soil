package fetch

import (
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"strings"
)

func List() []string {
	var result []string

	for _, c := range monitorConstant.Collectors {
		result = append(result, c.Name)
	}

	if s := environment.Optional(monitorConstant.PluginEnvironment); s != "" {
		if strings.HasPrefix(s, stringsConstant.Plus) {
			result = append(result, split.Comma(s)...)
		} else {
			result = split.Comma(s)
		}
	}

	return result
}
