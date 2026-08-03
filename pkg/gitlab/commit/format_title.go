package commit

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"strings"
)

func (c *Commit) formatTitle(f *option.Format) string {
	result := strings.TrimSpace(c.Title)

	if f.UseColor {
		return constant.Cyan("%s", result)
	}

	return result
}
