package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func Before() {
	a := argument.NewSimple("mattermost-before")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	f := constant.Format
	fmt.Printf("Channel: %s\n", c.Name)
	t := time.Now().Add(-30 * 24 * time.Hour)
	reference, found, e := m.FindPostBefore(c, t)

	if e != nil {
		fmt.Printf("Error: %s\n", e)

		return
	}

	if !found {
		fmt.Printf("No post before %s\n", t.Format(library.DateMinute))

		return
	}

	fmt.Printf("Reference: %s\n", reference.Format(f))
	fmt.Printf(
		"Date: %s\n",
		time.UnixMilli(reference.Raw.CreateAt).Format(library.DateMinute),
	)
	keep := 500
	posts := m.MustPostsBefore(c, t, keep)
	fmt.Printf(
		"Posts before %s or exceeding %d posts (%d found)\n",
		t.Format(library.DateMinute),
		keep,
		len(posts),
	)

	for _, p := range posts {
		fmt.Println(p.Format(f))
		fmt.Printf("  Time: %s\n", p.Create.Format(library.DateMinute))

		if false {
			m.MustDeletePost(p.Raw)
		}
	}
}
