package mock_indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustPush(
	name string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(name, body, metadata))
}
