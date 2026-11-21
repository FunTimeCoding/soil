package security_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
)

func (o *Tester) System() *system.Service {
	o.t.Helper()

	return system.New(
		o.Store(),
		o.Debian(),
		o.Score(),
		o.now,
		constant.StaleAfter,
		constant.TrackerStaleAfter,
	)
}
