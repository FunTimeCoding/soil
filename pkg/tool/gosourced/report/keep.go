package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"

func (r *Report) Keep(modules ...string) {
	var kept []*entry.Entry

	for _, e := range r.Entries {
		for _, m := range modules {
			if belongs(e.From, m) || belongs(e.To, m) {
				kept = append(kept, e)

				break
			}
		}
	}

	r.Entries = kept
}
