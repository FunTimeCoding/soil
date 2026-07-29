package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func Label() {
	c := confluence.NewEnvironment()
	f := constant.ConfluenceDense

	for _, l := range c.MustLabels() {
		fmt.Printf("Label: %+v\n", l)
	}

	c.SetLabels([]string{"favourite"})

	for _, o := range c.MustLabeled() {
		fmt.Println(o.Format(f))
	}
}
