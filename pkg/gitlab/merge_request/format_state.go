package merge_request

import (
	"github.com/funtimecoding/soil/pkg/console"
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
			return console.Yellow("%s", result)
		}

		if result == constant.ClosedState {
			return console.Green("%s", result)
		}
	}

	return result
}
