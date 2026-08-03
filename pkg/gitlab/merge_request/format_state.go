package merge_request

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (r *Request) formatState(f *option.Format) string {
	result := r.State

	if result == constant.OpenedState {
		result = constant.RequestOpenAlias
	}

	if f.UseColor {
		if result == constant.RequestOpenAlias {
			return consoleConstant.Yellow("%s", result)
		}

		if result == constant.ClosedState {
			return consoleConstant.Green("%s", result)
		}
	}

	return result
}
