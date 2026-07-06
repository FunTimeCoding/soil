package mock_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

func (i *Indexer) List(
	_ string,
	_ map[string]string,
	_ int,
	_ int,
	_ bool,
) (*face.ListOutcome, error) {
	return face.NewListOutcome(nil), nil
}
