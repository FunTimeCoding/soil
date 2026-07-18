package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func Labels(v []*label.Label) []*server.Label {
	result := make([]*server.Label, 0, len(v))

	for _, l := range v {
		result = append(result, Label(l))
	}

	return result
}
