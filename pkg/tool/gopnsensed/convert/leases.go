package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Leases(v []*lease.Lease) []server.Lease {
	result := []server.Lease{}

	for _, e := range v {
		result = append(result, *Lease(e))
	}

	return result
}
