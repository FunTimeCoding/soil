package change

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func state(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	actor string,
	at time.Time,
) []*event.Event {
	if after.State == before.State {
		return nil
	}

	return []*event.Event{
		event.New(constant.StateNoun, actor, StateLine(after.State, actor), at),
	}
}
