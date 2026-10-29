package gomattermost

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func formatThread(
	f *option.Format,
	t client.Thread,
) string {
	l := multi_line.New()
	l.Add(formatPost(f, t.Root))

	for _, p := range t.Replies {
		l.Add(strings.PrefixMultiline(formatPost(f, p), "> "))
	}

	return l.Render()
}
