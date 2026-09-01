package gowiki

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
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
		f.Tag(consoleConstant.TagCopyable)
	}

	if a.GetBoolean(argumentConstant.Watched) || a.GetBoolean(
		argumentConstant.Favorites,
	) {
		console.Line("Watch")

		for _, p := range c.MustWatched() {
			console.Line(p.Format(f))
			p.PrintConsole()
		}

		console.Line("Favorite")

		for _, p := range c.MustFavorites() {
			console.Line(p.Format(f))
		}

		return
	}

	for _, p := range c.MustPages() {
		console.Line(p.Format(f))
		p.PrintConsole()
	}
}
