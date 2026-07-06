package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
)

func Watch() {
	c := confluence.NewEnvironment()
	f := constant.Dense
	fmt.Println("Watch")

	for _, p := range c.MustWatched() {
		fmt.Println(p.Format(f))
	}

	fmt.Println("Favorite")

	for _, p := range c.MustFavorites() {
		fmt.Println(p.Format(f))
	}
}
