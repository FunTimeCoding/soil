package stream

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/types/stream_line"
	"path"
	"time"
)

func (c *Collector) record(b []byte) {
	var l stream_line.Line

	if json.Unmarshal(b, &l) != nil {
		return
	}

	if l.EventMessage == "" {
		return
	}

	t, e := time.Parse(constant.StreamTime, l.Timestamp)

	if e != nil {
		t = time.Now()
	}

	c.store.MustCreateEvent(
		event.Event{
			Time:      t,
			Process:   path.Base(l.ProcessImagePath),
			Subsystem: l.Subsystem,
			Category:  l.Category,
			Message:   l.EventMessage,
		},
	)
}
