package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/grafana"
)

func Read() {
	g := grafana.NewEnvironment()
	console.Line("Folders")

	for _, f := range g.Folders() {
		console.Format("  %s\n", f.Title)
	}

	console.Line("Dashboards")

	for _, d := range g.Dashboards() {
		console.Format("  %s\n", d.Title)
	}

	h := g.Home()
	console.Format("Home: %+v\n", h.Meta)
	console.Line("Search")

	for _, d := range g.Search() {
		console.Format("  %s\n", d.Title)
		console.Format("    %s\n", d.URL)
	}
}
