package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	f face.Reporter,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.EntriesTitle,
			constant.EntriesPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.AddEntryTitle,
			constant.AddEntryPath,
			webConstant.PaletteAction,
		),
	)

	return &Server{
		service:  v,
		reporter: f,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeArchive).
				WithStyle(
					join.Empty(webConstant.ComponentStyle, constant.InlineStyle),
				).
				WithLiveEndpoint(webConstant.LivePath).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
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
