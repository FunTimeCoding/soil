package memory_indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustPush(
	path string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(path, body, metadata))
}
