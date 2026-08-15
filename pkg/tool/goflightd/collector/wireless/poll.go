package wireless

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	flight "github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (c *Collector) poll() {
	r := run.New().NoPanic()
	r.Start(
		flight.Sudo,
		flight.SudoNonInteract,
		constant.Wdutil,
		constant.WdutilInformation,
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
				Kind:  flight.WirelessKind,
				Key:   key,
				Value: value,
			},
		)
	}
}
