package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func Space() {
	f := constant.ConfluenceFormat

	for _, s := range confluence.NewEnvironment().MustSpaces() {
		fmt.Println(s.Format(f))
	}
}
