package resolve

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func BuildFlags(directory string) []string {
	tags := build_tag.Discover(directory)

	if len(tags) == 0 {
		return nil
	}

	return []string{fmt.Sprintf("-tags=%s", join.Comma(tags))}
}
