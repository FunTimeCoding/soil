package change

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"slices"
	"time"
)

func approvals(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	at time.Time,
) []*event.Event {
	was := before.ApproverList()
	is := after.ApproverList()
	var result []*event.Event

	for _, name := range is {
		if !slices.Contains(was, name) {
			result = append(
				result,
				event.New(
					constant.ApprovalNoun,
					name,
					join.Empty(name, " approved"),
					at,
				),
			)
		}
	}

	for _, name := range was {
		if !slices.Contains(is, name) {
			result = append(
				result,
				event.New(
					constant.ApprovalNoun,
					name,
					join.Empty(name, " unapproved"),
					at,
				),
			)
		}
	}

	return result
}
