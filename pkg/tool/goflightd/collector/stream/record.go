package stream

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"path"
	"time"
)

func (c *Collector) record(b []byte) {
	var l line

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
