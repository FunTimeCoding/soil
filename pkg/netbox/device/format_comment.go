package device

import (
	"fmt"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/text"
)

func (d *Device) formatComment(f *option.Format) string {
	if d.Comment == constant.NoComment {
		return ""
	}

	o := text.OptimizeWhitespace(d.Comment, library.CompactWhitespace)

	if o == "" {
		return ""
	}

	var result string

	if strings.CountRune(o, '\n') > 0 {
		result = fmt.Sprintf(
			"\n%s",
			strings.Indent(o, 4, false),
		)
	} else {
		result = fmt.Sprintf(" %s", o)
	}

	if f.UseColor {
		result = consoleConstant.Magenta("%s", result)
	}

	return result
}
