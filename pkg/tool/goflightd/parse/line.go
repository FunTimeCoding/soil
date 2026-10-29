package parse

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/types/stream_line"
	"path"
	"time"
)

func Line(
	b []byte,
	now time.Time,
) *event.Event {
	var l stream_line.Line

	if json.Unmarshal(b, &l) != nil {
		return nil
	}

	if l.EventMessage == "" {
		return nil
	}

	t, e := time.Parse(constant.StreamTime, l.Timestamp)

	if e != nil {
		t = now
	}

	return event.New(
		t,
		path.Base(l.ProcessImagePath),
		l.Subsystem,
		l.Category,
		l.EventMessage,
	)
}
