package search_index

import "github.com/funtimecoding/soil/pkg/errors"

func (x *Index) Reindex(
	session string,
	path string,
) {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	_, e := x.remove(session)
	errors.PanicOnError(e)
	x.appendFile(session, path)
}
