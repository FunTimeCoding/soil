package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github"
)

func PullRequest() {
	g := github.NewEnvironment()
	u := g.MustUser()
	repositories := g.MustRepositories(u.Name)

	for _, r := range repositories {
		fmt.Printf("Repository: %s\n", *r.Name)

		for _, p := range g.MustPullRequests(u.Name, *r.Name) {
			fmt.Printf("  PR: %s\n", *p.Title)
			fmt.Printf("  %s\n", *p.HTMLURL)
		}
	}

	if true {
		for _, r := range repositories {
			fmt.Printf("Repository: %s\n", *r.Name)

			for _, i := range g.MustProjectIssues(u.Name, *r.Name) {
				fmt.Printf("  Issue: %s\n", *i.Title)
				fmt.Printf("  %s\n", *i.HTMLURL)
			}
		}
	}
}
