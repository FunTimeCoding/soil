package fetch

import (
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"strings"
)

func List() []string {
	var result []string

	for _, c := range item.Collectors {
		result = append(result, c.Name)
	}

	if s := environment.Optional(constant.PluginEnvironment); s != "" {
		if strings.HasPrefix(s, separator.Plus) {
			result = append(result, split.Comma(s)...)
		} else {
			result = split.Comma(s)
		}
	}

	return result
}
