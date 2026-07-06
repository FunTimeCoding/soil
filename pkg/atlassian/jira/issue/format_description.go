package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatDescription(f *option.Format) string {
	if i.Description == "" {
		return NoDescription
	}

	result := i.Description

	if f.UseColor {
		result = console.Magenta("%s", result)
	}

	return fmt.Sprintf("  Description: %s", result)
}
