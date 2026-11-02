package session_view

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/session_detail"

func New(d *session_detail.Detail) *View {
	return &View{Detail: d}
}
