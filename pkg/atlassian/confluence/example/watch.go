package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Watch() {
	c := confluence.NewEnvironment()
	f := constant.ConfluenceDense
	console.Line("Watch")

	for _, p := range c.MustWatched() {
		console.Line(p.Format(f))
	}

	console.Line("Favorite")

	for _, p := range c.MustFavorites() {
		console.Line(p.Format(f))
	}
}
