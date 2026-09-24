package report

import "github.com/funtimecoding/soil/pkg/crap/entry"

func (r *Report) Add(e *entry.Entry) {
	r.Entries = append(r.Entries, e)
}
