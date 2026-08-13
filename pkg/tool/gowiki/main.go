package gowiki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	console "github.com/funtimecoding/soil/pkg/console/constant"
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
		argumentConstant.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argumentConstant.Watched, false, "Watched")
	a.Boolean(argumentConstant.Favorites, false, "Favorites")
	a.Parse(version, gitHash, buildDate)
	c := confluence.NewEnvironment()
	f := atlassian.ConfluenceFormat.Copy()

	if a.GetBoolean(argumentConstant.Copyable) {
		f.Tag(console.TagCopyable)
	}

	if a.GetBoolean(argumentConstant.Watched) || a.GetBoolean(
		argumentConstant.Favorites,
	) {
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
