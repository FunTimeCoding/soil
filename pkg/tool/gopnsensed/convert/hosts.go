package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Hosts(v []*host.Host) []server.Host {
	result := []server.Host{}

	for _, e := range v {
		result = append(result, *Host(e))
	}

	return result
}
