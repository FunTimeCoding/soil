package pipeline_page

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/pipeline_view"

func New(v *pipeline_view.View) *Page {
	return &Page{view: v}
}
