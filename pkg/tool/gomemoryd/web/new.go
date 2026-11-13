package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
	f face.Reporter,
) *Server {
	search := palette.NewCommand(
		"Search memories",
		"/palette/memories",
		webConstant.PaletteSearch,
	)
	search.SwapTarget = ".palette-body"

	return &Server{
		service:  s,
		reporter: f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeCortex, favicon).
				WithStyle(constant.InlineStyle).
				WithItems(
					navigation_item.New(
						constant.MemoriesPath,
						constant.MemoriesTitle,
					),
					navigation_item.New(
						constant.RelationsPath,
						constant.RelationsTitle,
					),
					navigation_item.New(
						constant.ImpressionsPath,
						constant.ImpressionsTitle,
					),
					navigation_item.New(
						constant.StatisticPath,
						constant.StatisticTitle,
					),
					navigation_item.New(
						webConstant.SearchPath,
						constant.SearchTitle,
					),
				),
		),
	}
}
