package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func BuildInformation() {
	p := system.WorkDirectory()
	r := git.Open(p)
	h := git.Head(r).Hash()
	console.Format("Short hash: %s\n", h.String()[:8])

	if false {
		c := git.CommitFromHash(r, h)
		console.Format("Long hash: %s\n", h)
		console.Format("Author: %s\n", c.Author.Name)
		console.Format("Date: %s\n", c.Author.When)
		console.Format("Message: %s", c.Message)
	}

	latest := git.LatestTag(p)

	if latest == "" {
		console.Format("No tag found: %s\n", p)
	}

	console.Format("Latest: %s\n", latest)
	console.Format("Date: %s\n", time.Now().Format(constant.DateMinute))
}
