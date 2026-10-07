package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"path"
)

func kindStore(
	base string,
	k *kind.Kind,
) string {
	return path.Join(base, k.Name)
}
