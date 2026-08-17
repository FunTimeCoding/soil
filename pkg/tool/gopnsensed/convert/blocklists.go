package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Blocklists(v []*blocklist.Blocklist) []server.Blocklist {
	result := []server.Blocklist{}

	for _, e := range v {
		result = append(result, *Blocklist(e))
	}

	return result
}
