package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"time"
)

func (i *Issue) lastStatusTransition() time.Time {
	var result time.Time

	if i.Raw.Changelog == nil {
		return result
	}

	for _, h := range i.Raw.Changelog.Histories {
		for _, item := range h.Items {
			if item.Field == "status" {
				t, e := time.Parse(
					constant.TimeFormat,
					h.Created,
				)

				if e != nil {
					continue
				}

				if t.After(result) {
					result = t
				}
			}
		}
	}

	return result
}
