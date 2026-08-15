package bluetooth

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (c *Collector) poll() {
	r := run.New().NoPanic()
	r.Start(
		constant.SystemProfiler,
		constant.BluetoothSection,
		constant.ProfilerNotation,
	)

	if r.Error != nil {
		c.logger.Structured("system profiler failed", "error", r.Error.Error())

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
				Kind:  constant.BluetoothKind,
				Key:   key,
				Value: value,
			},
		)
	}
}
