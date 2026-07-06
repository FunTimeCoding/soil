package session

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/iterm/constant"
)

func (s *Session) formatTabTitle(f *option.Format) string {
	title := s.TabTitle

	if title == "" {
		title = constant.NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
