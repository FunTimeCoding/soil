package issue

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func (i *Issue) ChangeTime() time.Time {
	var result time.Time

	if u := i.LatestComment(); u != nil {
		result = CommentTime(u)
	} else {
		result = library.PickLatest(
			time.Time(i.Raw.Fields.Created),
			time.Time(i.Raw.Fields.Updated),
		)
	}

	if t := i.lastStatusTransition(); t.After(result) {
		result = t
	}

	return result
}
