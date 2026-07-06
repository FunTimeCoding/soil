package collect

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"time"
)

func tick(
	byName map[string]*LastRun,
	t time.Time,
) {
	for _, s := range constant.Collectors {
		if r, okay := byName[s.Name]; okay {
			if t.Sub(r.Time) >= s.Interval {
				r.Time = t
				runTime(s, t)
			}
		} else {
			byName[s.Name] = &LastRun{Command: s.Name, Time: t}
			runTime(s, t)
		}
	}
}
