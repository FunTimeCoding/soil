package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
)

func Changes(value any) []string {
	record, okay := value.(*run.Run)

	if !okay || record.Status != constant.StoreStatusSuccess {
		return nil
	}

	return ParseChanges(record.Output)
}
