package wireless

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (c *Collector) poll() {
	r := run.New().NoPanic()
	r.Start(
		constant.Sudo,
		constant.SudoNonInteract,
		systemConstant.Wdutil,
		systemConstant.WdutilInformation,
	)

	if r.Error != nil {
		c.logger.Structured("wdutil failed", "error", r.Error.Error())

		return
	}

	now := time.Now()

	for key, value := range Parse(r.OutputString) {
		if c.last[key] == value {
			continue
		}

		c.last[key] = value
		c.store.MustCreateSnapshot(
			snapshot.Snapshot{
				Time:  now,
				Kind:  constant.WirelessKind,
				Key:   key,
				Value: value,
			},
		)
	}
}
