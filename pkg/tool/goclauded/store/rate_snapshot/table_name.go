package rate_snapshot

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (Snapshot) TableName() string {
	return constant.RateSnapshotTable
}
