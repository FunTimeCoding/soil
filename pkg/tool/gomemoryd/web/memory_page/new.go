package memory_page

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_view"

func New(v *memory_view.View) *Page {
	return &Page{view: v}
}
