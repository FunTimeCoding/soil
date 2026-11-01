package session_page

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/session_view"

func New(v *session_view.View) *Page {
	return &Page{view: v}
}
