package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *store.Store) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.EntriesTitle,
			Path:     constant.EntriesPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.AddEntryTitle,
			Path:     constant.AddEntryPath,
			Category: web.PaletteAction,
		},
	)

	return &Server{
		store:    s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeArchive).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
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
