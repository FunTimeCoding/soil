package dashboard_page

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/dashboard_view"

func New(v *dashboard_view.View) *Page {
	return &Page{view: v}
}
