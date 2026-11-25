package change

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func Compare(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	actor string,
	at time.Time,
) []*event.Event {
	var result []*event.Event
	result = append(result, push(before, after, at)...)
	result = append(result, pipeline(before, after, at)...)
	result = append(result, approvals(before, after, at)...)
	result = append(result, draft(before, after, at)...)
	result = append(result, mergeability(before, after, at)...)

	return append(result, state(before, after, actor, at)...)
}
