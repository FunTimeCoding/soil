package search_result

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

type Result struct {
	Links []*link.Link
	Lists []*list.List
	Tags  []*tag.Tag
}
