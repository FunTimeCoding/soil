package server

import (
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func convertView(v *view.View) *server.View {
	result := &server.View{
		ViewId:   v.Identifier,
		WindowId: v.WindowIdentifier,
		FilePath: v.FilePath,
		Title:    v.Title,
		IsDirty:  v.Dirty,
		Syntax:   v.Syntax,
	}

	if v.Preview != "" {
		result.Preview = new(v.Preview)
	}

	if v.Text != "" {
		result.Text = new(v.Text)
	}

	return result
}
