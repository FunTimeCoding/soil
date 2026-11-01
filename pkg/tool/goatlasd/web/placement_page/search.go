package placement_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
)

func Search(name string) gomponents.Node {
	return entry.New(constant.PlacementsPath, constant.SearchAction).
		Query().
		Inline().
		Search(constant.SearchParameter, constant.SearchLabel, name).
		Placeholder(constant.SearchLabel).
		Focus().
		Node()
}
