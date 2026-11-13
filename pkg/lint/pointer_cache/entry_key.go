package pointer_cache

import "github.com/funtimecoding/soil/pkg/strings/join"

func entryKey(
	root string,
	document string,
) string {
	return join.Empty(root, "\x00", document)
}
