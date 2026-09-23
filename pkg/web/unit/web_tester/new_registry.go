package web_tester

import "github.com/funtimecoding/soil/pkg/web/palette"

func NewRegistry() *palette.Registry {
	r := palette.NewRegistry()
	r.Register(
		palette.Command{Label: "Dashboard", Path: "/", Category: "navigate"},
		palette.Command{
			Label:    "Create project",
			Path:     "/projects/new",
			Category: "action",
		},
		palette.Command{
			Label:    "Sessions",
			Path:     "/sessions",
			Category: "navigate",
		},
		palette.Command{
			Label:    "Start build",
			Path:     "/builds/start",
			Category: "action",
		},
		palette.Command{
			Label:    "Metrics",
			Path:     "/metrics",
			Category: "navigate",
		},
		palette.Command{
			Label:    "Push deploy",
			Path:     "/deploys/push",
			Category: "action",
		},
		palette.Command{
			Label:    "Search logs",
			Path:     "/logs/search",
			Category: "navigate",
		},
	)

	return r
}
