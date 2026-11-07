package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
	a *client.Client,
	f face.Reporter,
) *Server {
	return &Server{
		service:       s,
		authorization: a,
		reporter:      f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeAmethyst, favicon).
				WithStyle(constant.InlineStyle).
				WithItems(
					navigation_item.New(
						constant.ClientsPath,
						constant.ClientsTitle,
					),
					navigation_item.New(
						constant.CreatePath,
						constant.CreateTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
				),
		),
	}
}
