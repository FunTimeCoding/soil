package web_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
)

func NewRegistry() *registry.Registry {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			"Dashboard",
			constant.RootPath,
			constant.PaletteNavigate,
		),
		palette.NewCommand(
			"Create project",
			"/projects/new",
			constant.PaletteAction,
		),
		palette.NewCommand("Sessions", "/sessions", constant.PaletteNavigate),
		palette.NewCommand(
			"Start build",
			"/builds/start",
			constant.PaletteAction,
		),
		palette.NewCommand(
			"Metrics",
			constant.MetricsPath,
			constant.PaletteNavigate,
		),
		palette.NewCommand(
			"Push deploy",
			"/deploys/push",
			constant.PaletteAction,
		),
		palette.NewCommand(
			"Search logs",
			"/logs/search",
			constant.PaletteNavigate,
		),
	)

	return r
}
