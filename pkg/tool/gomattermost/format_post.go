package gomattermost

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/text"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func formatPost(
	f *option.Format,
	p client.Post,
) string {
	lines := split.NewLine(
		text.OptimizeWhitespace(p.Message, soil.CompactWhitespace),
	)

	for i, l := range lines {
		if l == "" {
			l = "(no text, image-only)"
		}

		if f.UseColor {
			l = constant.Cyan("%s", l)
		}

		lines[i] = l
	}

	s := status.New(f).String(
		p.CreatedAt.Format(timeConstant.HourMinute),
		p.Username,
		lines[0],
	)

	for _, l := range lines[1:] {
		s.Line("%s", l)
	}

	return s.Format()
}
