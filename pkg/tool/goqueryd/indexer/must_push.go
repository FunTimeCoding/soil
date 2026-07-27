package indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustPush(
	collection string,
	path string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(collection, path, body, metadata))
}
