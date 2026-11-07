package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	t face.Reporter,
) *Server {
	return &Server{
		service:  v,
		reporter: t,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeAmethyst, favicon).
				WithItems(
					navigation_item.New(
						constant.BookmarkPath,
						constant.BookmarkTitle,
					),
				),
		),
	}
}
