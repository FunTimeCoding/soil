package change

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func draft(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	at time.Time,
) []*event.Event {
	if after.Draft == before.Draft {
		return nil
	}

	line := "marked as ready"

	if after.Draft {
		line = "marked as draft"
	}

	return []*event.Event{event.New(constant.DraftNoun, "", line, at)}
}
