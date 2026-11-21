package overview

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
)

func fix(
	d *status_result.Result,
	unread error,
) string {
	if unread != nil || len(d.Statuses) == 0 {
		return constant.FixAbsent
	}

	for _, s := range d.Statuses {
		if s.Status != vulnerabilityConstant.DebianResolved {
			return constant.FixLocal
		}
	}

	return constant.FixDebian
}
