package resolve

import "github.com/funtimecoding/soil/pkg/source/build_tag"

func BuildFlags(directory string) []string {
	return TagFlags(build_tag.Discover(directory))
}
