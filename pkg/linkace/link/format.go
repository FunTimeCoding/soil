package link

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (l *Link) Format(f *option.Format) string {
	s := status.New(f).Integer(l.Identifier)

	for _, list := range l.Lists {
		s.DetailLink(
			fmt.Sprintf("https://%s/lists/%d", l.Host, list.Identifier),
			list.Name,
			"",
		)
	}

	for _, t := range l.Tags {
		s.DetailLink(
			fmt.Sprintf("https://%s/tags/%d", l.Host, t.Identifier),
			t.Name,
			"",
		)
	}

	title := l.Title

	if f.UseColor {
		title = constant.Cyan("%s", l.Title)
	}

	s.DetailLink(l.Link, title, "")

	if l.Description != "" {
		s.String(l.Description)
	}

	s.DetailLink(
		fmt.Sprintf("https://%s/links/%d", l.Host, l.Identifier),
		"Edit",
		"",
	)

	return s.Format()
}
