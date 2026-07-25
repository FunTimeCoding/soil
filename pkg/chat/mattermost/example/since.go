package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/errors"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func Since() {
	a := argument.NewSimple("mattermost-since")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	since := a.RequiredPositional(1, "SINCE")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	fmt.Printf("Channel: %s\n", c.Name)
	t, e := time.ParseInLocation(
		timeLibrary.DateMinute,
		since,
		time.Now().Location(),
	)
	errors.PanicOnError(e)
	posts := m.MustPostsSince(c, t)
	roots := 0
	replies := 0

	for _, p := range posts {
		if p.Raw.RootId == "" {
			roots++
		} else {
			replies++
		}
	}

	fmt.Printf(
		"Since %s: %d posts, %d roots, %d replies\n",
		since,
		len(posts),
		roots,
		replies,
	)

	for _, p := range posts {
		fmt.Printf(
			"  %s %s root=%v files=%d: %.60s\n",
			p.Create.Format(timeLibrary.DateMinute),
			p.User.Username,
			p.Raw.RootId == "",
			len(p.Raw.FileIds),
			p.Message,
		)
	}
}
