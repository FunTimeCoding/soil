package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/score_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ScoreLine(v *score_result.Result) gomponents.Node {
	text := fmt.Sprintf(
		"Scores: filling from the National Vulnerability Database, %d of %d vulnerabilities so far.",
		v.Start,
		v.Total,
	)

	if v.Filled {
		text = fmt.Sprintf(
			"Scores: %d vulnerabilities mirrored from the National Vulnerability Database, caught up %s.",
			v.Mirrored,
			time.FormatCompact(v.Synced),
		)
	}

	return html.P(html.Small(gomponents.Text(text)))
}
