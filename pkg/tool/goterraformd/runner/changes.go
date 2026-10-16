package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

func (r *Runner) changes(value any) *change.Change {
	record, okay := value.(*run.Run)

	if !okay || record.Status != constant.StoreStatusSuccess {
		return nil
	}

	return r.pending
}
