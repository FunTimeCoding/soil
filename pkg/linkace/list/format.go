package list

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (l *List) Format(f *option.Format) string {
	s := status.New(f).Integer(l.Identifier).DetailLink(
		fmt.Sprintf("https://%s/lists/%d", l.Host, l.Identifier),
		l.formatName(f),
		"",
	)

	if l.Description != "" {
		s.String(l.Description)
	}

	return s.Format()
}
