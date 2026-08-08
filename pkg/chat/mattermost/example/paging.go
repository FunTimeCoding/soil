package example

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func Paging() {
	a := argument.NewSimple("mattermost-paging")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	fmt.Printf("Channel: %s\n", c.Name)
	raw := model.NewAPIv4Client(
		locator.New(
			environment.Required(constant.MattermostHostEnvironment),
		).String(),
	)
	raw.SetOAuthToken(
		environment.Required(constant.MattermostTokenEnvironment),
	)
	background := context.Background()
	page, _, e := raw.GetPostsForChannel(
		background,
		c.Id,
		0,
		1000,
		constant.MattermostEmptyEntityTag,
		false,
		false,
	)
	errors.PanicOnError(e)
	fmt.Printf("Requested per_page 1000, received %d\n", len(page.Order))
	seen := map[string]int{}
	total := 0
	anchor := ""

	for {
		var chunk *model.PostList

		if anchor == "" {
			chunk, _, e = raw.GetPostsForChannel(
				background,
				c.Id,
				0,
				constant.MattermostMaxPerPage,
				constant.MattermostEmptyEntityTag,
				false,
				false,
			)
		} else {
			chunk, _, e = raw.GetPostsBefore(
				background,
				c.Id,
				anchor,
				0,
				constant.MattermostMaxPerPage,
				constant.MattermostEmptyEntityTag,
				false,
				false,
			)
		}

		errors.PanicOnError(e)

		if len(chunk.Order) == 0 {
			break
		}

		for _, identifier := range chunk.Order {
			seen[identifier]++
			total++
		}

		anchor = chunk.Order[len(chunk.Order)-1]
	}

	duplicate := 0

	for _, count := range seen {
		if count > 1 {
			duplicate++
		}
	}

	fmt.Printf(
		"Cursor walk: %d posts, %d unique, %d duplicated\n",
		total,
		len(seen),
		duplicate,
	)
}
