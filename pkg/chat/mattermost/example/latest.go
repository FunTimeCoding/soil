package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	chat "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func Latest() {
	a := argument.NewSimple("mattermost-latest")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	console.Format("Channel: %s\n", c.Name)
	limit := 10
	posts := m.MustLatestPosts(c, limit)
	console.Format("Latest %d posts (%d found)\n", limit, len(posts))

	for _, p := range posts {
		console.Line(p.Format(chat.MattermostFormat))
		console.Format("  Time: %s\n", p.Create.Format(constant.DateMinute))
	}
}
