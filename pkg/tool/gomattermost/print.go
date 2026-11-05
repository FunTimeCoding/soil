package gomattermost

import (
	"fmt"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/terminal"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost/thread_summary"
	"time"
)

func Print(
	t *terminal.Terminal,
	since time.Time,
	summaries []*thread_summary.Summary,
) {
	t.Line(
		fmt.Sprintf(
			"Unresolved threads since %s",
			since.Format(timeConstant.DateMinute),
		),
	)

	for _, s := range summaries {
		t.Line(formatThread(chatConstant.MattermostFormat, s.Thread))
		t.Line(consoleConstant.Magenta("%s", s.Summary))
	}

	if len(summaries) == 0 {
		t.Line("No unresolved threads")
	}
}
