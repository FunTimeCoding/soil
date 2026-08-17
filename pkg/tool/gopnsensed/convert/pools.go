package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Pools(v []*pool.Pool) []server.Pool {
	result := []server.Pool{}

	for _, e := range v {
		result = append(result, *Pool(e))
	}

	return result
}
