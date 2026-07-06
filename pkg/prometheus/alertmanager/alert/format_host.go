package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (a *Alert) formatHost(f *option.Format) string {
	result := a.Host()

	if result == "" {
		result = NoHost

		if f.UseColor {
			result = console.Yellow(result)
		}

		return result
	}

	if a.HostLink != "" &&
		!f.HasTag(tag.Copyable) &&
		!f.HasTag(tag.Markdown) {
		result = console.Link(a.HostLink, result, true)
	}

	return result
}
