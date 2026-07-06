package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/grafana"
)

func Read() {
	g := grafana.NewEnvironment()
	fmt.Println("Folders")

	for _, f := range g.Folders() {
		fmt.Printf("  %s\n", f.Title)
	}

	fmt.Println("Dashboards")

	for _, d := range g.Dashboards() {
		fmt.Printf("  %s\n", d.Title)
	}

	h := g.Home()
	fmt.Printf("Home: %+v\n", h.Meta)
	fmt.Println("Search")

	for _, d := range g.Search() {
		fmt.Printf("  %s\n", d.Title)
		fmt.Printf("    %s\n", d.URL)
	}
}
