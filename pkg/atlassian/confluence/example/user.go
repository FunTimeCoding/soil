package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func User() {
	fmt.Println(confluence.NewEnvironment().MustUser().Format(constant.ConfluenceFormat))
}
