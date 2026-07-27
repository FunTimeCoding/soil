package indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustDelete(collection string, path string) {
	errors.PanicOnError(i.Delete(collection, path))
}
