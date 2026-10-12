package wireless

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
)

func (c *Collector) Probe() bool {
	r := run.New().NoPanic()
	r.Start(
		constant.Sudo,
		constant.SudoNonInteract,
		systemConstant.Wdutil,
		systemConstant.WdutilInformation,
	)

	return r.Error == nil
}
