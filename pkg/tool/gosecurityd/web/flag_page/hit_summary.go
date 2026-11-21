package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
)

func hitSummary(c *affected_count.Count) string {
	if c == nil || c.Systems == 0 {
		return "No reporting system carries it."
	}

	return fmt.Sprintf(
		"%d systems: an upgrade fixes %d, Debian fix not in the candidate on %d, paid support only on %d, no fix yet on %d.",
		c.Systems,
		c.Upgrade,
		c.Pending,
		c.Paid,
		c.Open,
	)
}
