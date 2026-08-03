package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (a *Alert) formatHost(f *option.Format) string {
	result := a.Host()

	if result == "" {
		result = prometheus.NoHost

		if f.UseColor {
			result = constant.Yellow(result)
		}

		return result
	}

	if a.HostLink != "" &&
		!f.HasTag(constant.TagCopyable) &&
		!f.HasTag(constant.TagMarkdown) {
		result = console.Link(a.HostLink, result, true)
	}

	return result
}
