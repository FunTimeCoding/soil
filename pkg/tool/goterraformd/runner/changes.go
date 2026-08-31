package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/store"
)

func Changes(value any) []string {
	record, okay := value.(*store.Run)

	if !okay || record.Status != constant.StoreStatusSuccess {
		return nil
	}

	return ParseChanges(record.Output)
}
