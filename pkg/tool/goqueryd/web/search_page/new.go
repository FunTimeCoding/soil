package search_page

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func New(
	query string,
	collection string,
	mode string,
	metadata map[string]string,
) *Page {
	if mode != constant.KeywordMode {
		mode = constant.HybridMode
	}

	return &Page{
		query:      query,
		collection: collection,
		mode:       mode,
		metadata:   metadata,
	}
}
