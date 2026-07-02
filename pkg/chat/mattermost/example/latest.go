package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/time"
)

func Latest() {
	a := argument.NewSimple("mattermost-latest")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	fmt.Printf("Channel: %s\n", c.Name)
	limit := 10
	posts := m.MustLatestPosts(c, limit)
	fmt.Printf("Latest %d posts (%d found)\n", limit, len(posts))

	for _, p := range posts {
		fmt.Println(p.Format(constant.Format))
		fmt.Printf("  Time: %s\n", p.Create.Format(time.DateMinute))
	}
}
