package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	chat "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func Before() {
	a := argument.NewSimple("mattermost-before")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	f := chat.MattermostFormat
	console.Format("Channel: %s\n", c.Name)
	t := time.Now().Add(-30 * 24 * time.Hour)
	reference, found, e := m.FindPostBefore(c, t)

	if e != nil {
		console.Format("Error: %s\n", e)

		return
	}

	if !found {
		console.Format("No post before %s\n", t.Format(constant.DateMinute))

		return
	}

	console.Format("Reference: %s\n", reference.Format(f))
	console.Format(
		"Date: %s\n",
		time.UnixMilli(reference.Raw.CreateAt).Format(constant.DateMinute),
	)
	keep := 500
	posts := m.MustPostsBefore(c, t, keep)
	console.Format(
		"Posts before %s or exceeding %d posts (%d found)\n",
		t.Format(constant.DateMinute),
		keep,
		len(posts),
	)

	for _, p := range posts {
		console.Line(p.Format(f))
		console.Format("  Time: %s\n", p.Create.Format(constant.DateMinute))

		if false {
			m.MustDeletePost(p.Raw)
		}
	}
}
