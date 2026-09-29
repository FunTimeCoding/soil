package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

type SearchResult struct {
	Links []*link.Link
	Lists []*list.List
	Tags  []*tag.Tag
}
