package change

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func mergeability(
	before *snapshot.Snapshot,
	after *snapshot.Snapshot,
	at time.Time,
) []*event.Event {
	if before.MergeStatus == "" || after.MergeStatus == "" ||
		after.MergeStatus == before.MergeStatus {
		return nil
	}

	return []*event.Event{
		event.New(
			constant.MergeabilityNoun,
			"",
			fmt.Sprintf(
				"mergeability %s → %s",
				before.MergeStatus,
				after.MergeStatus,
			),
			at,
		),
	}
}
