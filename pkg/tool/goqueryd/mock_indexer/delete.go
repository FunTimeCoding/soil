package mock_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/call"

func (i *Indexer) Delete(
	collection string,
	path string,
) error {
	i.Deleted = append(i.Deleted, call.NewDelete(collection, path))

	return nil
}
