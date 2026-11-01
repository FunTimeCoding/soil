package collection_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"maragu.dev/gomponents"
)

func Pages(
	name string,
	sourceType string,
	page int,
	more bool,
) gomponents.Node {
	result := pager.New()

	if page > 1 {
		result.Previous(Path(name, sourceType, page-1), "← Newer")
	}

	if more {
		result.Next(Path(name, sourceType, page+1), "Older →")
	}

	return result.Node()
}
