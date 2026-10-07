package web

import (
	"github.com/funtimecoding/soil/pkg/raid_parser"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	elitePath string,
	outputPath string,
	p *raid_parser.Client,
	authorization *client.Client,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.LogsTitle,
			constant.LogsPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ReportsTitle,
			constant.ReportsPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.RaidsTitle,
			constant.RaidsPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.PlayersTitle,
			constant.PlayersPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(web.SignOutTitle, web.SignOutPath, web.PaletteAction),
	)

	return &Server{
		store:         s,
		elitePath:     elitePath,
		outputPath:    outputPath,
		parser:        p,
		authorization: authorization,
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeTyria).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(constant.LogsPath, constant.LogsTitle),
					navigation_item.New(
						constant.ReportsPath,
						constant.ReportsTitle,
					),
					navigation_item.New(
						constant.RaidsPath,
						constant.RaidsTitle,
					),
					navigation_item.New(
						constant.PlayersPath,
						constant.PlayersTitle,
					),
				),
		),
	}
}
