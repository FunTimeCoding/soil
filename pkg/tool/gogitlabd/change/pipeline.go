package change

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func pipeline(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	at time.Time,
) []*event.Event {
	if after.Pipeline == 0 {
		return nil
	}

	var line string

	switch {
	case after.Pipeline != before.Pipeline:
		line = fmt.Sprintf(
			"pipeline #%d %s",
			after.Pipeline,
			after.PipelineStatus,
		)
	case after.PipelineStatus != before.PipelineStatus:
		line = fmt.Sprintf(
			"pipeline #%d %s → %s",
			after.Pipeline,
			before.PipelineStatus,
			after.PipelineStatus,
		)
	default:
		return nil
	}

	return []*event.Event{event.New(constant.PipelineNoun, "", line, at)}
}
