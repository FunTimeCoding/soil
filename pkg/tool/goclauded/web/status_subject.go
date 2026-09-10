package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func statusSubject(subject string) gomponents.Node {
	if subject == "" {
		return gomponents.Text("")
	}

	if len(subject) != 36 {
		return gomponents.Text(subject)
	}

	return html.A(
		gomponents.Attr(
			"href",
			fmt.Sprintf("%s/%s", constant.SessionsPath, subject),
		),
		html.Small(gomponents.Text(subject)),
	)
}
