package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (a *Alert) formatHost(f *option.Format) string {
	result := a.Host()

	if result == "" {
		result = constant.NoHost

		if f.UseColor {
			result = consoleConstant.Yellow(result)
		}

		return result
	}

	if a.HostLink != "" &&
		!f.HasTag(consoleConstant.TagCopyable) &&
		!f.HasTag(consoleConstant.TagMarkdown) {
		result = console.Link(a.HostLink, result, true)
	}

	return result
}
