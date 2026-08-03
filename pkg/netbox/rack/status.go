package rack

import (
	"fmt"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (r *Rack) status(f *option.Format) string {
	var result string

	if r.Raw.Status != nil {
		if *r.Raw.Status.Label == constant.RackActiveLabel &&
			*r.Raw.Status.Value == constant.RackActive {
			result = constant.RackActive
		} else if *r.Raw.Status.Label == constant.RackDeprecatedLabel &&
			*r.Raw.Status.Value == constant.RackDeprecated {
			result = constant.RackDeprecated

			if f.UseColor {
				result = consoleConstant.Yellow("%s", result)
			}
		} else {
			result = fmt.Sprintf(
				"%s (%s=%s)",
				constant.RackUnexpected,
				*r.Raw.Status.Label,
				*r.Raw.Status.Value,
			)

			if f.UseColor {
				result = consoleConstant.Red("%s", result)
			}
		}
	} else {
		result = constant.RackUnknown

		if f.UseColor {
			result = consoleConstant.Red("%s", result)
		}
	}

	return result
}
