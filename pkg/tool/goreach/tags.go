package goreach

import (
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Tags(root string) string {
	return join.Comma(build_tag.Discover(root))
}
