package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func BuildInformation() {
	p := system.WorkDirectory()
	r := git.Open(p)
	h := git.Head(r).Hash()
	fmt.Printf("Short hash: %s\n", h.String()[:8])

	if false {
		c := git.CommitFromHash(r, h)
		fmt.Printf("Long hash: %s\n", h)
		fmt.Printf("Author: %s\n", c.Author.Name)
		fmt.Printf("Date: %s\n", c.Author.When)
		fmt.Printf("Message: %s", c.Message)
	}

	latest := git.LatestTag(p)

	if latest == "" {
		fmt.Printf("No tag found: %s\n", p)
	}

	fmt.Printf("Latest: %s\n", latest)
	fmt.Printf("Date: %s\n", time.Now().Format(library.DateMinute))
}
