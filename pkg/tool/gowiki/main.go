package gowiki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	wiki "github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gowiki/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argument.Watched, false, "Watched")
	a.Boolean(argument.Favorites, false, "Favorites")
	a.Parse(version, gitHash, buildDate)
	c := confluence.NewEnvironment()
	f := wiki.Format.Copy()

	if a.GetBoolean(argument.Copyable) {
		f.Tag(tag.Copyable)
	}

	if a.GetBoolean(argument.Watched) || a.GetBoolean(argument.Favorites) {
		fmt.Println("Watch")

		for _, p := range c.MustWatched() {
			fmt.Println(p.Format(f))
			p.PrintConsole()
		}

		fmt.Println("Favorite")

		for _, p := range c.MustFavorites() {
			fmt.Println(p.Format(f))
		}

		return
	}

	for _, p := range c.MustPages() {
		fmt.Println(p.Format(f))
		p.PrintConsole()
	}
}
