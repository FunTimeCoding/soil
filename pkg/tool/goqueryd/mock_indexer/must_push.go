package mock_indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustPush(
	collection string,
	name string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(collection, name, body, metadata))
}
