package usage_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
	"time"
)

func Note(
	result *usage_result.Result,
	c *credential.Credential,
) gomponents.Node {
	var items []string

	if result != nil {
		items = append(
			items,
			fmt.Sprintf("Updated %s", library.FormatCompact(result.LastUpdated)),
		)
	}

	if c != nil && c.Expired(time.Now()) {
		items = append(
			items,
			fmt.Sprintf(
				"credential expired %s, start a Claude Code session to refresh",
				library.FormatCompact(c.ExpiresAt),
			),
		)
	}

	if len(items) == 0 {
		return gomponents.Group(nil)
	}

	return html.P(
		html.Class(webConstant.AbsentClass),
		gomponents.Text(strings.Join(items, " - ")),
	)
}
