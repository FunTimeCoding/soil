package runner

import "github.com/funtimecoding/soil/pkg/strings/join"

func (r *Runner) formatTags() string {
	if len(r.Tags) == 0 {
		return ""
	}

	return join.Comma(r.Tags)
}
