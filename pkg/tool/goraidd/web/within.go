package web

import (
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"time"
)

func within(
	all []fight.Fight,
	start string,
	end string,
) []fight.Fight {
	var result []fight.Fight

	for _, f := range all {
		if start != "" {
			t, e := time.Parse(constant.FilterFormat, start)

			if e == nil && f.Timestamp.Before(t) {
				continue
			}
		}

		if end != "" {
			t, e := time.Parse(constant.FilterFormat, end)

			if e == nil && f.Timestamp.After(t) {
				continue
			}
		}

		result = append(result, f)
	}

	return result
}
