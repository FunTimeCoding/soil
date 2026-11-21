package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
)

func AffectedCount(v *affected_count.Count) server.AffectedCount {
	return server.AffectedCount{
		Systems: v.Systems,
		Reboot:  v.Reboot,
		Upgrade: v.Upgrade,
		Pending: v.Pending,
		Paid:    v.Paid,
		Open:    v.Open,
	}
}
