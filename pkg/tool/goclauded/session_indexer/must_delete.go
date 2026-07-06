package session_indexer

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Indexer) MustDelete(path string) {
	errors.PanicOnError(i.Delete(path))
}
