package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/summary"
)

func Clients(v []*client.Client) []*summary.Summary {
	result := make([]*summary.Summary, 0, len(v))

	for _, c := range v {
		result = append(result, Client(c))
	}

	return result
}
