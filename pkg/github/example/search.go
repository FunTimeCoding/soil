package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/github"
)

func Search() {
	c := github.NewEnvironment()

	if false {
		for _, r := range c.MustSearchRepository("user:%s", c.MustUser().Name) {
			console.Format("Repository: %s\n", r.Name)
		}
	}

	if true {
		for _, r := range c.ActionRepository() {
			console.Format("Code: %+v\n", *r.Raw.Name)
		}
	}
}
