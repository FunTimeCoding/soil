package web

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func published(r *record.Record) string {
	if r.Published == nil {
		return "pending"
	}

	return "yes"
}
