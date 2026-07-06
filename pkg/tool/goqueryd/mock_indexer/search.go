package mock_indexer

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
)

func (i *Indexer) Search(_ *search_option.Option) ([]face.SearchResult, error) {
	return nil, nil
}
