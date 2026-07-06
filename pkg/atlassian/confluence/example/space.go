package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
)

func Space() {
	f := constant.Format

	for _, s := range confluence.NewEnvironment().MustSpaces() {
		fmt.Println(s.Format(f))
	}
}
