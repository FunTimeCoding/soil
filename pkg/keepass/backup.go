package keepass

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"time"
)

func (c *Client) Backup(at time.Time) string {
	target := fmt.Sprintf("%s.%s.bak", c.path, at.Format("2006-01-02-150405"))
	system.CopyFile(c.path, target)

	return target
}
