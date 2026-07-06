package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
)

func User() {
	fmt.Println(confluence.NewEnvironment().MustUser().Format(constant.Format))
}
