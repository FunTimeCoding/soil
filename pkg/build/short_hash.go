package build

import "github.com/funtimecoding/soil/pkg/git/constant"

func ShortHash(hash string) string {
	if len(hash) < constant.HashLength {
		return hash
	}

	return hash[:constant.HashLength]
}
