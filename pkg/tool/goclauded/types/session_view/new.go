package session_view

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/session_detail"

func New(d *session_detail.Detail) *View {
	return &View{Detail: d}
}
