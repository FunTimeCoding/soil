package git

import "github.com/funtimecoding/soil/pkg/git/constant"

func ShortHash(path string) string {
	return Head(Open(path)).Hash().String()[:constant.HashLength]
}
