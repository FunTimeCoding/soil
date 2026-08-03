package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatDescription(f *option.Format) string {
	if i.Description == "" {
		return constant.JiraNoDescription
	}

	result := i.Description

	if f.UseColor {
		result = consoleConstant.Magenta("%s", result)
	}

	return fmt.Sprintf("  Description: %s", result)
}
