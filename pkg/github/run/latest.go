package run

import "github.com/funtimecoding/soil/pkg/github/constant"

func Latest(v []*Run) *Run {
	var latest *Run

	for _, r := range v {
		if r.Status != constant.CompletedStatus {
			continue
		}

		if latest == nil {
			latest = r

			continue
		}

		if r.Create.After(latest.Create) {
			latest = r
		}
	}

	return latest
}
