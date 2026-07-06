package status

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
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

	if s.format.HasTag(tag.Copyable) {
		return s.TagLine(tag.Copyable, "  %s", prefixed(prefix, link))
	}

	if s.format.HasTag(tag.Markdown) {
		return s.TagLine(
			tag.Markdown,
			"  %s",
			prefixed(prefix, markdown.Link(label, link)),
		)
	}

	s.bubbles = append(s.bubbles, console.Link(link, label, true))

	return s
}
