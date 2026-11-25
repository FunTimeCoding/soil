package change

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func push(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	at time.Time,
) []*event.Event {
	if before.Commit == "" || after.Commit == before.Commit {
		return nil
	}

	short := after.Commit[:min(len(after.Commit), constant.CommitPrefix)]

	return []*event.Event{
		event.New(constant.PushNoun, "", join.Empty("pushed ", short), at),
	}
}
