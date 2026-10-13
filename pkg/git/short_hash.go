package git

import "github.com/funtimecoding/soil/pkg/git/constant"

func ShortHash(path string) string {
	return Head(MustOpen(path)).Hash().String()[:constant.HashLength]
}
