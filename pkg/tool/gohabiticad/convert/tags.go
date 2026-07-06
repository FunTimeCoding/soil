package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func Tags(v []*tag.Tag) []*server.Tag {
	result := make([]*server.Tag, 0, len(v))

	for _, t := range v {
		result = append(result, Tag(t))
	}

	return result
}
