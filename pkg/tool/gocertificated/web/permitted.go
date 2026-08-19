package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func permitted(r *record.Record) string {
	c := r.Material().Certificate
	var result []string
	result = append(result, c.PermittedDNSDomains...)

	for _, a := range c.PermittedIPRanges {
		result = append(result, a.String())
	}

	return join.CommaSpace(result)
}
