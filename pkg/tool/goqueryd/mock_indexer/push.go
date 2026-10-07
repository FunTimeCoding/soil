package mock_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/call"

func (i *Indexer) Push(
	collection string,
	name string,
	body string,
	metadata map[string][]string,
) error {
	i.Pushed = append(i.Pushed, call.NewPush(collection, name, body, metadata))

	return nil
}
