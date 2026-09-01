package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/github"
)

func PullRequest() {
	g := github.NewEnvironment()
	u := g.MustUser()
	repositories := g.MustRepositories(u.Name)

	for _, r := range repositories {
		console.Format("Repository: %s\n", *r.Name)

		for _, p := range g.MustPullRequests(u.Name, *r.Name) {
			console.Format("  PR: %s\n", *p.Title)
			console.Format("  %s\n", *p.HTMLURL)
		}
	}

	if true {
		for _, r := range repositories {
			console.Format("Repository: %s\n", *r.Name)

			for _, i := range g.MustProjectIssues(u.Name, *r.Name) {
				console.Format("  Issue: %s\n", *i.Title)
				console.Format("  %s\n", *i.HTMLURL)
			}
		}
	}
}
