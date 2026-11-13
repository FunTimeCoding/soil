package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	n face.EventNotifier,
	f face.Reporter,
) *Server {
	return &Server{
		service:  v,
		notifier: n,
		reporter: f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeArchive, favicon).
				WithStyle(constant.InlineStyle).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.EntriesPath,
						constant.EntriesTitle,
					),
					navigation_item.New(
						constant.AddEntryPath,
						constant.AddEntryTitle,
					),
				),
		),
	}
}
