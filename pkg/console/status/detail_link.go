package status

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown"
)

func (s *Status) DetailLink(
	link string,
	label string,
	prefix string,
) *Status {
	if link == "" {
		return s
	}

	if label == "" {
		label = "Link"
	}

	if s.format.HasTag(constant.TagCopyable) {
		return s.TagLine(constant.TagCopyable, "  %s", prefixed(prefix, link))
	}

	if s.format.HasTag(constant.TagMarkdown) {
		return s.TagLine(
			constant.TagMarkdown,
			"  %s",
			prefixed(prefix, markdown.Link(label, link)),
		)
	}

	s.bubbles = append(s.bubbles, console.Link(link, label, true))

	return s
}
