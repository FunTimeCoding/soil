package wireless

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	flight "github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
)

func (c *Collector) Probe() bool {
	r := run.New().NoPanic()
	r.Start(
		flight.Sudo,
		flight.SudoNonInteract,
		constant.Wdutil,
		constant.WdutilInformation,
	)

	return r.Error == nil
}
