package file_report

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (r *Report) Add(c *concern.Concern) {
	r.Concerns = append(r.Concerns, c)
}
